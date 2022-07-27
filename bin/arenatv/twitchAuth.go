package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type twitchAuth struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type streamData struct {
	Data []struct {
		ID           string    `json:"id"`
		UserID       string    `json:"user_id"`
		UserLogin    string    `json:"user_login"`
		UserName     string    `json:"user_name"`
		GameID       string    `json:"game_id"`
		GameName     string    `json:"game_name"`
		Type         string    `json:"type"`
		Title        string    `json:"title"`
		ViewerCount  int       `json:"viewer_count"`
		StartedAt    time.Time `json:"started_at"`
		Language     string    `json:"language"`
		ThumbnailURL string    `json:"thumbnail_url"`
		TagIds       []string  `json:"tag_ids"`
		IsMature     bool      `json:"is_mature"`
		isOnline     bool
	} `json:"data"`
	Pagination struct {
	} `json:"pagination"`
}

func getTwitchToken() <-chan twitchAuth {
	url := "https://id.twitch.tv/oauth2/token"
	method := "POST"
	payload := strings.NewReader("client_id=3f7ouqiwbe2x7mpqsl3qvipx4kjkxm&client_secret=3uc4xzm9vncape8va8eq7wkwcmddug&grant_type=client_credentials")
	client := &http.Client{}
	println("Called getTwitchToken...")
	tt := make(chan twitchAuth)
	go func() {
		req, err := http.NewRequest(method, url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			app.Log(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			app.Log(err)
			return
		}
		var twitchToken twitchAuth
		err = json.Unmarshal(b, &twitchToken)
		if err != nil {
			panic(err)
		}
		tt <- twitchAuth(twitchToken)
	}()
	return tt
}

func getMultipleStreamsData(streams []liveStream, token string) <-chan streamData {
	request := ""
	for _, streamer := range streams {
		request += "user_login=" + streamer.Slug + "&"
	}

	url := "https://api.twitch.tv/helix/streams?" + request
	method := "GET"
	payload := strings.NewReader("")
	client := &http.Client{}
	println("Called getStreamData...")

	sd := make(chan streamData)

	go func() {
		req, err := http.NewRequest(method, url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Client-Id", "3f7ouqiwbe2x7mpqsl3qvipx4kjkxm")
		req.Header.Add("Authorization", "Bearer "+token)
		if err != nil {
			app.Log(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			app.Log(err)
			return
		}
		var streamdata streamData

		err = json.Unmarshal(b, &streamdata)
		if err != nil {
			panic(err)
		}

		sd <- streamData(streamdata)
	}()
	return sd
}

func getStreamData(user_name string, token string) <-chan streamData {
	url := "https://api.twitch.tv/helix/streams?user_login=" + user_name
	method := "GET"
	payload := strings.NewReader("")
	client := &http.Client{}
	println("Called getStreamData...")

	sd := make(chan streamData)

	go func() {
		req, err := http.NewRequest(method, url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Client-Id", "3f7ouqiwbe2x7mpqsl3qvipx4kjkxm")
		req.Header.Add("Authorization", "Bearer "+token)
		if err != nil {
			app.Log(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			app.Log(err)
			return
		}
		var streamdata streamData

		err = json.Unmarshal(b, &streamdata)
		if err != nil {
			panic(err)
		}
		sd <- streamData(streamdata)
	}()
	return sd
}
