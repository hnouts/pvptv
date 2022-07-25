package main

import (
	"path"
	"sort"
	"strings"
	"time"
)

type liveStream struct {
	Slug      string
	MainClass string
	ClassList []string
	Name      string
	URL       string
	Viewers   int
	Online    bool
	Title     string
	Cards     []string
	Links     []socialLink
	AddedAt   time.Time
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
			Slug:      "snupy",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			Name:  "Snupy",
			URL:   "https://www.twitch.tv/snupy",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/snupytv",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC8epYaippj44sS41P3nLfkw/",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/h2yAPBB",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/Snupy/",
				},
			},
		},
		{
			Slug:      "bicmexwow",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name:  "Bicmexwow",
			URL:   "https://www.twitch.tv/bicmexwow",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/bicmexwow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/NullesWa",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/6Rdaz8K",
				},
			},
		},
		{
			Slug:      "venruki",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name:  "Venruki",
			URL:   "https://www.twitch.tv/venruki",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/ElliottVenczel",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/venruki",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/venczel/",
				},
			},
		},
		{
			Slug:      "raikubest",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name:  "Raikubest",
			URL:   "https://www.twitch.tv/raikubest",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Raiku_Wow",
				},
			},
		},
		{
			Slug:      "xaryu",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name:  "Xaryu",
			URL:   "https://www.twitch.tv/xaryu",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Raiku_Wow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/xaryu?sub_confirmation=1",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/joshlujan/",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/xaryu",
				},
				{
					Slug: "website",
					URL:  "https://xaryu.tv/",
				},
			},
		},
		{
			Slug:      "hydramist",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name:  "Hydramist",
			URL:   "https://www.twitch.tv/hydramist",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/hydramist",
				},
				{
					Slug: "youtube",
					URL:  "https://youtube.com/hydramist",
				},
				{
					Slug: "facebook",
					URL:  "http://www.facebook.com/pages/Hydramist/147296701974820?ref=hl",
				},
			},
		},
		{
			Slug:      "zenlyn",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name:  "Zenlyn",
			URL:   "https://www.twitch.tv/zenlyn",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:      "anboniwow",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name:  "Anboniwow",
			URL:   "https://www.twitch.tv/anboniwow",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://www.twitter.com/anboniwow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/anboniwow",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/EhKKHEm",
				},
			},
		},
		{
			Slug:      "chastv",
			MainClass: "priest",
			ClassList: []string{
				"priest",
				"druid",
			},
			Name:  "Chastv",
			URL:   "https://www.twitch.tv/chastv",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:      "accident_",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name:  "Accident",
			URL:   "https://www.twitch.tv/accident_",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:      "pika_pala",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name:  "Pikachu",
			URL:   "https://www.twitch.tv/pika_pala",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/pala_pika",
				},
			},
		},
		{
			Slug:      "mirlolxd",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name:  "Mirlol",
			URL:   "https://www.twitch.tv/mirlolxd",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/mirlolxd",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC3r7qQGEPcRuHGmMNCZdS4Q",
				},
			},
		},
		{
			Slug:      "palumor",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name:  "Palumor",
			URL:   "https://www.twitch.tv/palumor",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/iPalumor",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/Palumor",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/ydJ5eUpZWD",
				},
			},
		},
		{
			Slug:      "drainerx",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
				"monk",
				"druid",
				"priest",
				"paladin",
			},
			Name:  "Drainerx",
			URL:   "https://www.twitch.tv/drainerx",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/intent/user?screen_name=Drainerxtv",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/drainerxtv?sub_confirmation=1",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/DMUSWhG",
				},
			},
		},
		{
			Slug:      "zhreytv",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name:  "Zhreytv",
			URL:   "https://www.twitch.tv/zhreytv",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Zhreytv",
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
