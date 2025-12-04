package twitch

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

// HelixClient is a minimal Twitch Helix client with in-memory caching for Get Streams.
type HelixClient struct {
	client       *http.Client
	clientID     string
	clientSecret string

	mu          sync.Mutex
	accessToken string
	tokenExpiry time.Time

	streamsCache     map[string]streamsCacheEntry // key: sorted, comma-joined logins
	streamsCacheTTL  time.Duration
}

type streamsCacheEntry struct {
	data HelixStreamsResponse
	exp  time.Time
}

type HelixStreamsResponse struct {
	Data []struct {
		ID          string    `json:"id"`
		UserID      string    `json:"user_id"`
		UserLogin   string    `json:"user_login"`
		UserName    string    `json:"user_name"`
		GameID      string    `json:"game_id"`
		GameName    string    `json:"game_name"`
		Type        string    `json:"type"`
		Title       string    `json:"title"`
		ViewerCount int       `json:"viewer_count"`
		StartedAt   time.Time `json:"started_at"`
		Language    string    `json:"language"`
		Thumbnail   string    `json:"thumbnail_url"`
		IsMature    bool      `json:"is_mature"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// NewHelixClient constructs a HelixClient using env vars for credentials.
//  - TWITCH_CLIENT_ID
//  - TWITCH_CLIENT_SECRET
func NewHelixClient() (*HelixClient, error) {
	id := os.Getenv("TWITCH_CLIENT_ID")
	secret := os.Getenv("TWITCH_CLIENT_SECRET")
	if id == "" || secret == "" {
		return nil, errors.New("TWITCH_CLIENT_ID and TWITCH_CLIENT_SECRET must be set")
	}

	return &HelixClient{
		client:        &http.Client{Timeout: 5 * time.Second},
		clientID:      id,
		clientSecret:  secret,
		streamsCache:  make(map[string]streamsCacheEntry),
		streamsCacheTTL: 60 * time.Second,
	}, nil
}

// GetStreamsByLogin calls Helix Get Streams for user_logins with 60s cache.
// Handles batching automatically if more than 100 logins are provided.
func (c *HelixClient) GetStreamsByLogin(ctx context.Context, logins []string) (HelixStreamsResponse, error) {
	if len(logins) == 0 {
		return HelixStreamsResponse{}, nil
	}

	// Normalize and sort key for cache.
	norm := make([]string, 0, len(logins))
	for _, l := range logins {
		if l == "" {
			continue
		}
		norm = append(norm, strings.ToLower(l))
	}
	cacheKey := strings.Join(norm, ",")

	now := time.Now()

	// Fast path: cache
	c.mu.Lock()
	if entry, ok := c.streamsCache[cacheKey]; ok && entry.exp.After(now) {
		c.mu.Unlock()
		return entry.data, nil
	}
	c.mu.Unlock()

	// Ensure access token
	if err := c.ensureToken(ctx); err != nil {
		return HelixStreamsResponse{}, err
	}

	// Build request
	// Twitch API allows up to 100 user_login parameters per request
	// and up to 100 results per request (via first parameter)
	// We need to batch if we have more than 100 logins
	const maxLoginsPerRequest = 100
	
	var allStreams []struct {
		ID          string    `json:"id"`
		UserID      string    `json:"user_id"`
		UserLogin   string    `json:"user_login"`
		UserName    string    `json:"user_name"`
		GameID      string    `json:"game_id"`
		GameName    string    `json:"game_name"`
		Type        string    `json:"type"`
		Title       string    `json:"title"`
		ViewerCount int       `json:"viewer_count"`
		StartedAt   time.Time `json:"started_at"`
		Language    string    `json:"language"`
		Thumbnail   string    `json:"thumbnail_url"`
		IsMature    bool      `json:"is_mature"`
	}
	
	// Process in batches of 100
	for i := 0; i < len(norm); i += maxLoginsPerRequest {
		end := i + maxLoginsPerRequest
		if end > len(norm) {
			end = len(norm)
		}
		batch := norm[i:end]
		
		u, _ := url.Parse("https://api.twitch.tv/helix/streams")
		q := u.Query()
		q.Set("first", "100") // Request up to 100 results
		for _, login := range batch {
			q.Add("user_login", login)
		}
		u.RawQuery = q.Encode()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
		if err != nil {
			return HelixStreamsResponse{}, err
		}
		req.Header.Set("Client-Id", c.clientID)
		req.Header.Set("Authorization", "Bearer "+c.accessToken)

		resp, err := c.client.Do(req)
		if err != nil {
			return HelixStreamsResponse{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return HelixStreamsResponse{}, errors.New("twitch helix streams request failed: " + resp.Status)
		}

		var batchResp HelixStreamsResponse
		if err := json.NewDecoder(resp.Body).Decode(&batchResp); err != nil {
			return HelixStreamsResponse{}, err
		}
		
		// Collect streams from this batch
		// When querying by user_login, Twitch returns all matching streams for those users
		// (up to the 'first' limit, which we set to 100)
		allStreams = append(allStreams, batchResp.Data...)
	}
	
	// Combine all results
	parsed := HelixStreamsResponse{
		Data: allStreams,
	}

	// Store in cache
	c.mu.Lock()
	c.streamsCache[cacheKey] = streamsCacheEntry{
		data: parsed,
		exp:  now.Add(c.streamsCacheTTL),
	}
	c.mu.Unlock()

	return parsed, nil
}

func (c *HelixClient) ensureToken(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.accessToken != "" && time.Now().Before(c.tokenExpiry.Add(-30*time.Second)) {
		return nil
	}

	form := url.Values{}
	form.Set("client_id", c.clientID)
	form.Set("client_secret", c.clientSecret)
	form.Set("grant_type", "client_credentials")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://id.twitch.tv/oauth2/token", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("twitch token request failed: " + resp.Status)
	}

	var body struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}

	c.accessToken = body.AccessToken
	c.tokenExpiry = time.Now().Add(time.Duration(body.ExpiresIn) * time.Second)
	return nil
}


