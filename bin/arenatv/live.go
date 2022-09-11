package main

import (
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
	Links     []socialLink
	AddedAt   time.Time
}

type socialLink struct {
	Slug string
	URL  string
}

func (r liveStream) twitchSlug() string {
	return r.Slug
}

func getLiveStreamersForAClass(class string) []liveStream {
	defer timer()()
	streams := getLiveStreamers()
	println(class)
	var filteredStreams []liveStream

	for _, streamer := range streams {
		if isCurrentClass(streamer, class) { // create new array of lives for current class
			filteredStreams = append(filteredStreams, streamer) // enrich lives with viewer count and online status
		}
	}
	return filteredStreams
}

func isCurrentClass(streamer liveStream, class string) bool {
	for _, b := range streamer.ClassList {
		if b == class {
			return true
		}
	}
	return false
}

func getLiveStreamers() []liveStream {
	streams := []liveStream{
		//DH
		{
			Slug:      "fuseton",
			MainClass: "demon_hunter",
			ClassList: []string{
				"demon_hunter",
				"warrior",
				"monk",
				"paladin",
				"rogue",
			},
			Name: "Fuseton",
			URL:  "https://www.twitch.tv/fuseton",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/fuseton",
				},
			},
		},
		{
			Slug:      "exzistancetw",
			MainClass: "demon_hunter",
			ClassList: []string{
				"demon_hunter",
				"warrior",
			},
			Name: "Exzistancetw",
			URL:  "https://www.twitch.tv/exzistancetw",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/whocareslul",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCQU6XlJgVKMDRv3NnqgdIJw",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/BvHMFpY",
				},
			},
		},
		{
			Slug:      "mvqq",
			MainClass: "demon_hunter",
			ClassList: []string{
				"demon_hunter",
			},
			Name: "Mvqq",
			URL:  "https://www.twitch.tv/mvqq",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Mvqdh",
				},
				{
					Slug: "youtube",
					URL:  "https://youtube.com/c/mvqdh",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/andreas_bergeer/",
				},
			},
		},
		{
			Slug:      "trenacetatetv",
			MainClass: "demon_hunter",
			ClassList: []string{
				"demon_hunter",
			},
			Name: "Trenacetate",
			URL:  "https://www.twitch.tv/trenacetatetv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/trenacetatetv",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCZwnwgvQxNX99RqfnBjlifw",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/agqdg2A",
				},
				{
					Slug: "instagram",
					URL:  "http://instagram.com/jamaldinsworld/",
				},
			},
		},
		//DK
		{
			Slug:      "skippeetv",
			MainClass: "death_knight",
			ClassList: []string{
				"death_knight",
			},
			Name: "Skippee",
			URL:  "https://www.twitch.tv/skippeetv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/skippeedk",
				},
			},
		},
		{
			Slug:      "exyth1",
			MainClass: "death_knight",
			ClassList: []string{
				"death_knight",
			},
			Name:  "Exyth",
			URL:   "https://www.twitch.tv/exyth1",
			Links: []socialLink{},
		},
		{
			Slug:      "notmes",
			MainClass: "death_knight",
			ClassList: []string{
				"death_knight",
				"warrior",
			},
			Name: "Mes",
			URL:  "https://www.twitch.tv/notmes",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/mes_wow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC3yooAg-uF1xwa1OHQ1NK6g",
				},
			},
		},
		{
			Slug:      "bobydk1",
			MainClass: "death_knight",
			ClassList: []string{
				"death_knight",
			},
			Name: "Bobydk1",
			URL:  "https://www.twitch.tv/bobydk1",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/bobydk1",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/FMmxmzZkQe",
				},
			},
		},
		//DRUID
		{
			Slug:      "claak",
			MainClass: "druid",
			ClassList: []string{
				"druid",
				"priest",
			},
			Name: "Claak",
			URL:  "https://www.twitch.tv/claak",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/claak_tv",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCPESiEiyz59_KfY6vemeo4g",
				},
			},
		},
		{
			Slug:      "tonyferalmovies",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			Name: "Tonyferalmovies",
			URL:  "https://www.twitch.tv/tonyferalmovies",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/tonyferalmovies",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/Tonyferalmovies?sub_confirmation=1",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/443Xe3U",
				},
			},
		},
		{
			Slug:      "healingstat",
			MainClass: "druid",
			ClassList: []string{
				"druid",
				"paladin",
				"priest",
				"monk",
			},
			Name: "Healingstat",
			URL:  "https://www.twitch.tv/healingstat",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Healingstat",
				},
			},
		},
		{
			Slug:      "luuxia",
			MainClass: "druid",
			ClassList: []string{
				"druid",
				"paladin",
				"priest",
			},
			Name: "Luuxia",
			URL:  "https://www.twitch.tv/luuxia",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/JLuuxiA",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/xtKy3pk",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/luuxiaoff/?hl=fr",
				},
			},
		},
		{
			Slug:      "sodapoppin",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			Name: "Sodapoppin",
			URL:  "https://www.twitch.tv/sodapoppin",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Sodapoppintv/",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/sodapoppin",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/sodapoppin",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/sodapoppintv/",
				},
			},
		},
		{
			Slug:      "snupy",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			Name: "Snupy",
			URL:  "https://www.twitch.tv/snupy",
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
			Slug:      "pszjager",
			MainClass: "druid",
			ClassList: []string{
				"druid",
				"priest",
				"monk",
			},
			Name: "Pszjager",
			URL:  "https://www.twitch.tv/pszjager",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/pszjager",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCL-IFyMkjFeMU9Aiee1f8DQ",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/hxDy7FU",
				},
			},
		},
		{
			Slug:      "haraw",
			MainClass: "druid",
			ClassList: []string{
				"druid",
				"demon_hunter",
				"warrior",
			},
			Name: "Haraw",
			URL:  "https://www.twitch.tv/haraw",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Harawlul",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/HARAWW",
				},
			},
		},
		{
			Slug:      "ilovelucy_",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			Name: "Ilovelucy",
			URL:  "https://www.twitch.tv/ilovelucy_",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/ilovelucy_wow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/playlist?list=PLg2w4zWZQhvc6opGac6_3OxqxyGQiSOuy",
				},
			},
		},
		{
			Slug:      "supatease",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			Name: "Supatease",
			URL:  "https://www.twitch.tv/supatease",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Supatease",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCt6INY6VJraOZBVWTJKmvkQ",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/supacasting/",
				},
			},
		},
		{
			Slug:      "minpojke",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			Name: "Minpojke",
			URL:  "https://www.twitch.tv/minpojke",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Minpojke_",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/minpojke",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/STTWtwwDet",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/jonandersson92/",
				},
			},
		},
		//HUNTER
		{
			Slug:      "paradise1_",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name: "Paradise1",
			URL:  "https://www.twitch.tv/paradise1_",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Paradise1WoW",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC3Ui8dwMLAY_GMrFqj5ztEw?",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/NR38xs7",
				},
			},
		},
		{
			Slug:      "Disyo",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name: "Disyo",
			URL:  "https://www.twitch.tv/Disyo",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Disyowow",
				},
			},
		},
		{
			Slug:      "kasuxoxo",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name: "Kasuxoxo",
			URL:  "https://www.twitch.tv/kasuxoxo",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/kasuxoxoz",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCu5LyP1luajA3c8EZ6kgI7g?view_as=subscriber",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/kasuhoho/",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/5W3YWV7Hbn",
				},
			},
		},
		{
			Slug:      "billybobbs",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name: "Billybobbs",
			URL:  "https://www.twitch.tv/billybobbs",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC7evJCfohlRtzNSY5QU3Mcg",
				},
			},
		},
		{
			Slug:      "dilly_wow",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name:  "Dilly",
			URL:   "https://www.twitch.tv/dilly_wow",
			Links: []socialLink{},
		},
		{
			Slug:      "ssds",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name: "Ssds",
			URL:  "https://www.twitch.tv/ssds",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/SsdsLw",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC_6s9qDs2NZme0Rnc9rN2_g",
				},
			},
		},
		{
			Slug:      "jellybeans",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name: "Jellybeans",
			URL:  "https://www.twitch.tv/jellybeans",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/JellybeansTV",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCsfnT_8Dr5j3xpvz0azHmaA?view_as=subscriber",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/vincenttrangg/",
				},
			},
		},
		{
			Slug:      "bicmexwow",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name: "Bicmexwow",
			URL:  "https://www.twitch.tv/bicmexwow",
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
			Slug:      "freezion",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			Name: "Freezion",
			URL:  "https://www.twitch.tv/freezion",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/watch?v=gdGHz0GVXH8",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/gtuMn3f",
				},
			},
		},
		//MAGE
		{
			Slug:      "cmdxyz",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Cmdxyz",
			URL:  "https://www.twitch.tv/cmdxyz",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/cmdzyx",
				},
			},
		},
		{
			Slug:      "pherix1",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Pherix1",
			URL:  "https://www.twitch.tv/pherix1",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/DG_pherix",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/hXVUb2",
				},
			},
		},
		{
			Slug:      "vultz",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Vultz",
			URL:  "https://www.twitch.tv/vultz",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Vultzover",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/XTXgsRx",
				},
			},
		},
		{
			Slug:      "verinasty",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name:  "Verinasty",
			URL:   "https://www.twitch.tv/verinasty",
			Links: []socialLink{},
		},
		{
			Slug:      "samiyam",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Samiyam",
			URL:  "https://www.twitch.tv/samiyam",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/samiyamwow",
				},
			},
		},
		{
			Slug:      "wealthymanx",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Wealthyman",
			URL:  "https://www.twitch.tv/wealthymanx",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/C9Wealthyman",
				},
			},
		},
		{
			Slug:      "wowalec",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Alec",
			URL:  "https://www.twitch.tv/wowalec",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCBXOjQDF5Q96Qf0KNttBqow",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/TempoAlec",
				},
			},
		},
		{
			Slug:      "hansol",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Hansol",
			URL:  "https://www.twitch.tv/hansol",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "http://www.youtube.com/azncoolz",
				},
				{
					Slug: "twitter",
					URL:  "http://www.twitter.com/hansolonfire",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/vk82Gj6",
				},
			},
		},
		{
			Slug:      "draizn",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Draizn",
			URL:  "https://www.twitch.tv/draizn",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC5b9NGeDXyBTg-RVjNhm47A",
				},
			},
		},
		{
			Slug:      "albraik",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Albraik",
			URL:  "https://www.twitch.tv/albraik",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Albraiik",
				},
			},
		},
		{
			Slug:      "venruki",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Venruki",
			URL:  "https://www.twitch.tv/venruki",
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
			Slug:      "ziqoftw",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Ziqoftw",
			URL:  "https://www.twitch.tv/ziqoftw",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://www.twitter.com/ziqoftw",
				},
				{
					Slug: "youtube",
					URL:  "http://www.youtube.com/Ziqoftw",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/ziqoftw/",
				},
			},
		},
		{
			Slug:      "zqitv",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Zqitv",
			URL:  "https://www.twitch.tv/zqitv",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/Muffintoppinlolz/videos?view_as=subscriber",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/tarz6n/",
				},
			},
		},
		{
			Slug:      "raikubest",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			Name: "Raikubest",
			URL:  "https://www.twitch.tv/raikubest",
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
			Name: "Xaryu",
			URL:  "https://www.twitch.tv/xaryu",
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
		//MONK
		{
			Slug:      "banwellx",
			MainClass: "monk",
			ClassList: []string{
				"monk",
			},
			Name: "Banwellx",
			URL:  "https://www.twitch.tv/banwellx",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://www.twitter.com/banwell",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/banwellx",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/WtFZvrA",
				},
			},
		},
		{
			Slug:      "woopy",
			MainClass: "monk",
			ClassList: []string{
				"monk",
			},
			Name: "Woopy",
			URL:  "https://www.twitch.tv/woopy",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Woopywow",
				},
			},
		},
		{
			Slug:      "chunliww",
			MainClass: "monk",
			ClassList: []string{
				"monk",
			},
			Name:  "Chunliww",
			URL:   "https://www.twitch.tv/chunliww",
			Links: []socialLink{},
		},
		{
			Slug:      "trilltko",
			MainClass: "monk",
			ClassList: []string{
				"monk",
				"demon_hunter",
				"rogue",
			},
			Name: "Trilltko",
			URL:  "https://www.twitch.tv/trilltko",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/TrilltkoWW",
				},
			},
		},
		{
			Slug:      "orbrexth",
			MainClass: "monk",
			ClassList: []string{
				"monk",
				"shaman",
			},
			Name: "Orbrexth",
			URL:  "https://www.twitch.tv/orbrexth",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/orbreXth",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/orbreXth/videos",
				},
			},
		},
		{
			Slug:      "trillebartom",
			MainClass: "monk",
			ClassList: []string{
				"warrior",
				"rogue",
				"priest",
			},
			Name: "Trillebartom",
			URL:  "https://www.twitch.tv/trillebartom",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/trillebartom",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/trillebartom?sub_confirmation=1",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/trille",
				},
			},
		},
		//PRIEST
		{
			Slug:      "hozitojones",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name: "Hozitojones",
			URL:  "https://www.twitch.tv/hozitojones",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Hozitojones",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC1cd0NFWO8XxgLhg23F1y-w?view_as=subscriber",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/sebas_toriello/?hl=es-la",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/V5eNAW5",
				},
			},
		},
		{
			Slug:      "krawnzlol",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name: "Krawnzlol",
			URL:  "https://www.twitch.tv/krawnzlol",
			Links: []socialLink{
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/jraranda7",
				},
			},
		},
		{
			Slug:      "wizkx",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name: "Wizkx",
			URL:  "https://www.twitch.tv/wizkx",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/wizkxd",
				},
			},
		},
		{
			Slug:      "vilaye",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name:  "Vilaye",
			URL:   "https://www.twitch.tv/vilaye",
			Links: []socialLink{},
		},
		{
			Slug:      "hydramist",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name: "Hydramist",
			URL:  "https://www.twitch.tv/hydramist",
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
			Links: []socialLink{},
		},
		{
			Slug:      "oddwaffletbc",
			MainClass: "priest",
			ClassList: []string{
				"priest",
				"warrior",
			},
			Name: "Oddwaffle",
			URL:  "https://www.twitch.tv/oddwaffletbc",
			Links: []socialLink{
				{
					Slug: "mail",
					URL:  "oddwaffle1@gmail.com",
				},
			},
		},
		{
			Slug:      "anboniwow",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name: "Anboniwow",
			URL:  "https://www.twitch.tv/anboniwow",
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
			Slug:      "mehhx",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name: "Mehhx",
			URL:  "https://www.twitch.tv/mehhx",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/JuhaniHalme",
				},
			},
		},
		{
			Slug:      "gekz",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			Name: "Gekz",
			URL:  "https://www.twitch.tv/gekz",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Gekzs",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/6M6S8RBhKR",
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
			Links: []socialLink{},
		},
		{
			Slug:      "Eltoni52",
			MainClass: "priest",
			ClassList: []string{
				"priest",
				"druid",
			},
			Name: "Eltoni52",
			URL:  "https://www.twitch.tv/Eltoni52",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/eltoni25",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCPwd-YmkC1aZGX7Gk91ucZw",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/fcxQfrYe5r",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/eltoni522/",
				},
			},
		},
		//PALADIN
		{
			Slug:      "abxtv",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
				"priest",
			},
			Name:  "abxtv",
			URL:   "https://www.twitch.tv/abxtv",
			Links: []socialLink{},
		},
		{
			Slug:      "borngood",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name: "Borngood",
			URL:  "https://www.twitch.tv/borngood",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/BorngoodW",
				},
			},
		},
		{
			Slug:      "rubcub",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name: "Rubcub",
			URL:  "https://www.twitch.tv/rubcub",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/PG_Rubcub",
				},
			},
		},
		{
			Slug:      "savix",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name: "Savix",
			URL:  "https://www.twitch.tv/savix",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/SavixIrL",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCIRe2YighqgPmSDjSb3Fpiw?view_as=subscriber",
				},
				{
					Slug: "reddit",
					URL:  "https://www.reddit.com/r/Savix/",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/savix",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/savixirl/",
				},
			},
		},
		{
			Slug:      "pvplab",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name: "Pvplab",
			URL:  "https://www.twitch.tv/pvplab",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "http://twitter.com/pvplablive",
				},
				{
					Slug: "discord",
					URL:  "http://discord.gg/pvplab",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/pvplablive/",
				},
			},
		},
		{
			Slug:      "accident_",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name:  "Accident",
			URL:   "https://www.twitch.tv/accident_",
			Links: []socialLink{},
		},
		{
			Slug:      "bushy25",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name:  "Bushy25",
			URL:   "https://www.twitch.tv/bushy25",
			Links: []socialLink{},
		},
		{
			Slug:      "tintinlives",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name:  "Tintinlives",
			URL:   "https://www.twitch.tv/tintinlives",
			Links: []socialLink{},
		},
		{
			Slug:      "snowmixy",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
				"priest",
			},
			Name: "Snowmixy",
			URL:  "https://www.twitch.tv/snowmixy",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Snowmixy",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/Snowmixy",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/snowmixy/",
				},
			},
		},
		{
			Slug:      "vanguardstv",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name: "Vanguards",
			URL:  "https://www.twitch.tv/vanguardstv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/VanguardsTV",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/Vanguards",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/Vanguards",
				},
			},
		},
		{
			Slug:      "pika_pala",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name: "Pikachu",
			URL:  "https://www.twitch.tv/pika_pala",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/pala_pika",
				},
			},
		},
		{
			Slug:      "zabius",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			Name: "Zabius",
			URL:  "https://www.twitch.tv/zabius",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/FelixFahauer1",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCZSfktks-roK7sdh6DmdgMQ",
				},
			},
		},
		//ROGUE
		{
			Slug:      "traxelol",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name:  "Traxelol",
			URL:   "https://www.twitch.tv/traxelol",
			Links: []socialLink{},
		},
		{
			Slug:      "jiberjaber",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Jiberjaber",
			URL:  "https://www.twitch.tv/jiberjaber",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/PfqC5gP7Mp",
				},
			},
		},
		{
			Slug:      "chickchau",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Chickchau",
			URL:  "https://www.twitch.tv/chickchau",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/jX7rfQHUZX",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/chick_chau/",
				},
			},
		},
		{
			Slug:      "inquisx",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Inquisx",
			URL:  "https://www.twitch.tv/inquisx",
			Links: []socialLink{
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/kushkaboss",
				},
			},
		},
		{
			Slug:      "NAIIKAII",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Naiikaii",
			URL:  "https://www.twitch.tv/NAIIKAII",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/NaiikaiiGG",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/NAIIKAII?sub_confirmation=1",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/eEMUwuS",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/naiikaii_/",
				},
			},
		},
		{
			Slug:      "nahj",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Nahj",
			URL:  "https://www.twitch.tv/nahj",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/RogueNahj",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/nahjwow",
				},
			},
		},
		{
			Slug:      "mirlolxd",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Mirlol",
			URL:  "https://www.twitch.tv/mirlolxd",
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
			Name: "Palumor",
			URL:  "https://www.twitch.tv/palumor",
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
			Slug:      "pikabooirl",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Pikaboo",
			URL:  "https://www.twitch.tv/pikabooirl",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/pikaboowow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCrYo6np138Ge6iCRMgS93aQ?view_as=subscriber",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/jasonsmithtm/",
				},
			},
		},
		{
			Slug:      "stungodx",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
				"priest",
			},
			Name: "Stungodx",
			URL:  "https://www.twitch.tv/stungodx",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Stungodx",
				},
			},
		},
		{
			Slug:      "aidenzx",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Aidenzx",
			URL:  "https://www.twitch.tv/aidenzx",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/UN_Aiden",
				},
			},
		},
		{
			Slug:      "SensusLive",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Sensuslive",
			URL:  "https://www.twitch.tv/SensusLive",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/SensusYT",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/SensusYT",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/jzJ3P6eEGr",
				},
			},
		},
		{
			Slug:      "jaxington",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Jaxington",
			URL:  "https://www.twitch.tv/jaxington",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/jaxirl",
				},
				{
					Slug: "youtube",
					URL:  "https://youtube.com/skidsh",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/jExCyhv",
				},
			},
		},
		{
			Slug:      "akrololz",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Akrololz",
			URL:  "https://www.twitch.tv/akrololz",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Akrololz",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/akrololz?sub_confirmation=1",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/akrololz",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/akrobro/",
				},
			},
		},
		{
			Slug:      "whaazz",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Whaazz",
			URL:  "https://www.twitch.tv/whaazz",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Whaazz",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCJFHadK6prEPYQuCmR0tsug",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/oscar_wulff/",
				},
			},
		},
		{
			Slug:      "psherotv",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Pshero",
			URL:  "https://www.twitch.tv/psherotv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Psherotv",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/psheor",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/h7nrY8j",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/psherotv/",
				},
			},
		},
		{
			Slug:      "knoffwow",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
				"monk",
				"warrior",
			},
			Name: "Knoffwow",
			URL:  "https://www.twitch.tv/knoffwow",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://www.twitter.com/knoffwow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/knoffwow",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/atKfD52hfm",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/knoffwow/",
				},
			},
		},
		{
			Slug:      "frozentherogue",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			Name: "Frozen",
			URL:  "https://www.twitch.tv/frozentherogue",
			Links: []socialLink{
				{
					Slug: "website",
					URL:  "https://www.donationalerts.com/r/frozentherogue",
				},
			},
		},
		//SHAMAN
		{
			Slug:      "warbla",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			Name:  "Warbla",
			URL:   "https://www.twitch.tv/warbla",
			Links: []socialLink{},
		},
		{
			Slug:      "kolowavex",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
				"monk",
			},
			Name: "Kolowavex",
			URL:  "https://www.twitch.tv/kolowavex",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Kolohots",
				},
			},
		},
		{
			Slug:      "tiqqlethis",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			Name: "Tiqqlethis",
			URL:  "https://www.twitch.tv/tiqqlethis",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/TiqqleThis",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/tiqqle?sub_confirmation=1",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/gsky4Ykcg8",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/tiqqleofficial/",
				},
			},
		},
		{
			Slug:      "swapxy",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			Name: "Swapxy",
			URL:  "https://www.twitch.tv/swapxy",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Swapxy",
				},
			},
		},
		{
			Slug:      "thesterge",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			Name: "Thesterge",
			URL:  "https://www.twitch.tv/thesterge",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/stergey",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/thesolrac5",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/carloscorreagg/",
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
			Name: "Drainerx",
			URL:  "https://www.twitch.tv/drainerx",
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
			Slug:      "zeepeye",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			Name: "Zeepeye",
			URL:  "https://www.twitch.tv/zeepeye",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/ZeepeyeSwE",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCmRgHhywwS7fp97tki9eY6g",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/tCq9TQGaQm",
				},
			},
		},
		{
			Slug:      "lontartv",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
				"druid",
			},
			Name: "Lontar",
			URL:  "https://www.twitch.tv/lontartv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Lontar_wow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCCDFA0pRgHYZPx3l_uCRvAw",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/gabri_cano/",
				},
			},
		},
		{
			Slug:      "cdewx",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
				"druid",
				"priest",
				"paladin",
			},
			Name: "Cdew",
			URL:  "https://www.twitch.tv/cdewx",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/cdew_wow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/tenderloinx1?sub_confirmation=1",
				},
			},
		},
		{
			Slug:      "Lovelesstv",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
				"rogue",
				"monk",
				"demon_hunter",
			},
			Name: "Loveless",
			URL:  "https://www.twitch.tv/Lovelesstv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/lovelesswow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCsLbmQM0iqdM5_Fg2jRh3Gg",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/eFrG6fmM",
				},
			},
		},
		//WARLOCK
		{
			Slug:      "bualock",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name: "Bualock",
			URL:  "https://www.twitch.tv/bualock",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/BuaLockOW",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/bualock/videos",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/sxhwe3b",
				},
				{
					Slug: "instagram",
					URL:  "https://instagram.com/BuaLockOW",
				},
			},
		},
		{
			Slug:      "snutzy",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name: "Snutzy",
			URL:  "https://www.twitch.tv/snutzy",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Snutztv",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/kelnguyenn/",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/eMTkkB3",
				},
			},
		},
		{
			Slug:      "wallirik",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name:  "Wallirik",
			URL:   "https://www.twitch.tv/wallirik",
			Links: []socialLink{},
		},
		{
			Slug:      "jazggz",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name: "Jazggz",
			URL:  "https://www.twitch.tv/jazggz",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/jazggz",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCXWELAuYNuRk53qgIWPTFaA",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/ZfNg2n44WK",
				},
			},
		},
		{
			Slug:      "anniefuchsia",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
				"paladin",
				"priest",
			},
			Name: "Anniefuchsia",
			URL:  "https://www.twitch.tv/anniefuchsia",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/anniefuchsia",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/anniefuchsia?sub_confirmation=1",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/anniefuchsia/",
				},
			},
		},
		{
			Slug:      "maldiva",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name: "Maldiva",
			URL:  "https://www.twitch.tv/maldiva",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Maldivawow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/Maldivapvp",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/maldiva",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/maldivawow/",
				},
			},
		},
		{
			Slug:      "chanimaly",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name: "Chanimaly",
			URL:  "https://www.twitch.tv/chanimaly",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/c9_chan",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/ChanimalTV/",
				},
			},
		},
		{
			Slug:      "germinouu066",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name: "Germi",
			URL:  "https://www.twitch.tv/germinouu066",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Germinou06",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/RHDBxsN",
				},
			},
		},
		{
			Slug:      "featherfeeet",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name:  "Featherfeeet",
			URL:   "https://www.twitch.tv/featherfeeet",
			Links: []socialLink{},
		},
		{
			Slug:      "gelubabatv",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name: "Gelubabatv",
			URL:  "https://www.twitch.tv/gelubabatv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/CalGelu",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCskmJY6W20-v27tuIr5eDkA",
				},
			},
		},
		{
			Slug:      "mercedesa",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
				"death_knight",
			},
			Name: "Mercedesa",
			URL:  "https://www.twitch.tv/mercedesa",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/mercewow",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/Mau5mWQ",
				},
			},
		},
		{
			Slug:      "dakkroth",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			Name: "Dakkroth",
			URL:  "https://www.twitch.tv/dakkroth",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Dakkrothwow",
				},
				{
					Slug: "facebook",
					URL:  "https://www.facebook.com/Dakkroth",
				},
			},
		},
		//WARRIOR
		{
			Slug:      "zhreytv",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Zhreytv",
			URL:  "https://www.twitch.tv/zhreytv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Zhreytv",
				},
			},
		},
		{
			Slug:      "blizo",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Blizo",
			URL:  "https://www.twitch.tv/blizo",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/blizowow",
				},
			},
		},
		{
			Slug:      "smexxin",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Smexxin",
			URL:  "https://www.twitch.tv/smexxin",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Smexxin",
				},
			},
		},
		{
			Slug:      "mitu_tv",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Mitu",
			URL:  "https://www.twitch.tv/mitu_tv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Mitu_TV",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC85rl8ND2SNxO3LC_tCTztw?view_as=subscriber",
				},
			},
		},
		{
			Slug:      "tay_warr",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Tay",
			URL:  "https://www.twitch.tv/tay_warr",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Tay39826924",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC0aMJOAQZ1oAQSUBH6kA-Fg/videos",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/thgrty1/",
				},
			},
		},
		{
			Slug:      "hito",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
				"druid",
			},
			Name: "Hito",
			URL:  "https://www.twitch.tv/hito",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Hitowow",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/Hitokiro",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/CNj5bjnBB8",
				},
			},
		},
		{
			Slug:      "joefernandes123",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Joefernandes",
			URL:  "https://www.twitch.tv/joefernandes123",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/joefernandes123",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCRU4UWnsYB-zhaKkaF37EEw",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/MMpmYK3F",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/joefernandes123/",
				},
			},
		},
		{
			Slug:      "bajheera",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Bajheera",
			URL:  "https://www.twitch.tv/bajheera",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/BajheeraWoW",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/BajheeraWoW",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/DWc3mFmEyd",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/bajheerawow/",
				},
			},
		},
		{
			Slug:      "dekel",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Dekel",
			URL:  "https://www.twitch.tv/dekel",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/dekeldeclan",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCMEwlq2-4FzgzeGllkRHJSA?view_as=subscriber",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/jPb9tr9",
				},
				{
					Slug: "instagram",
					URL:  "https://instagram.com/dekeldeclan",
				},
			},
		},
		{
			Slug:      "magnusz",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Magnusz",
			URL:  "https://www.twitch.tv/magnusz",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/sx6_magnusz",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/magnusz",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/dFtqrHe",
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
