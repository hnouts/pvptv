package main

import (
	"path"
	"sort"
	"strings"
	"time"
)

type liveStream struct {
	Slug    string
	Class   string
	Name    string
	URL     string
	Viewers int
	Online  bool
	Title   string
	Cards   []string
	Links   []socialLink
	AddedAt time.Time
}

type socialLink struct {
	Slug string
	URL  string
}

func (r liveStream) twitchID() string {
	return path.Base(r.URL)
}

func getLiveStreamers() []liveStream {
	streams := []liveStream{
		{
			Slug:  "hydramist",
			Class: "priest",
			Name:  "Hydramist",
			URL:   "https://www.twitch.tv/hydramist",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5yx6BWlEVcY",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
		{
			Slug:    "zenlyn",
			Class:   "priest",
			Name:    "Zenlyn",
			URL:     "https://www.twitch.tv/zenlyn",
			Viewers: 0,
			Online:  false,
			Cards:   []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5yx6BWlEVcY",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
		{
			Slug:    "anboniwow",
			Class:   "priest",
			Name:    "Anboniwow",
			URL:     "https://www.twitch.tv/anboniwow",
			Viewers: 0,
			Online:  false,
			Cards:   []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5yx6BWlEVcY",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
		{
			Slug:    "accident_",
			Class:   "paladin",
			Name:    "Accident",
			URL:     "https://www.twitch.tv/accident_",
			Viewers: 0,
			Online:  false,
			Cards:   []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5yx6BWlEVcY",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
		{
			Slug:    "mirlolxd",
			Class:   "rogue",
			Name:    "Mirlol",
			URL:     "https://www.twitch.tv/mirlolxd",
			Viewers: 0,
			Online:  false,
			Cards:   []string{},
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://chillhop.com",
				},
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5yx6BWlEVcY",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/chillhop",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chillhopmusic",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/groups/1561371024098016",
				},
			},
		},
	}

	sort.Slice(streams, func(a, b int) bool {
		return strings.Compare(streams[a].Name, streams[b].Name) < 0
	})

	for _, r := range streams {
		sort.Slice(r.Links, func(a, b int) bool {
			return strings.Compare(r.Links[a].Slug, r.Links[b].Slug) < 0
		})
	}

	return streams
}

func socialIcon(slug string) string {
	switch slug {
	case "youtube":
		return youtubeSVG

	case "reddit":
		return redditSVG

	case "facebook":
		return facebookSVG

	case "instagram":
		return instagramSVG

	case "twitter":
		return twitterSVG

	case "discord":
		return discordSVG

	case "website":
		return websiteSVG

	default:
		return linkSVG
	}
}
