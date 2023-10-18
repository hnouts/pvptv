package main

import (
	"time"
)

type liveStream struct {
	Slug      string
	MainClass string
	ClassList []string
	SpecList  []specList
	Name      string
	URL       string
	Viewers   int
	GameName  string
	Online    bool
	Title     string
	Links     []socialLink
	AddedAt   time.Time
}

type specList struct {
	SPClass string
	SPSpec  string
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
			Slug:      "volimond",
			MainClass: "demon_hunter",
			ClassList: []string{
				"demon_hunter",
			},
			SpecList: []specList{
				{
					SPClass: "demon_hunter",
					SPSpec:  "vengeance",
				},
			},
			Name: "Volimond",
			URL:  "https://www.twitch.tv/volimond",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Tank4daze",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/irontriangle",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/volimond/",
				},
			},
		},
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
			SpecList: []specList{
				{
					SPClass: "demon_hunter",
					SPSpec:  "havoc",
				},
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
			SpecList: []specList{
				{
					SPClass: "demon_hunter",
					SPSpec:  "havoc",
				},
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
			Slug:      "ltrainetv",
			MainClass: "death_knight",
			ClassList: []string{
				"priest",
				"death_knight",
			},
			Name: "Ltrainetv",
			URL:  "https://www.twitch.tv/ltrainetv",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/@LtraineTV",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/ltrainetv/",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/7V7C6Qjd",
				},
			},
		},
		{
			Slug:      "razoonxz",
			MainClass: "death_knight",
			ClassList: []string{
				"death_knight",
			},
			SpecList: []specList{
				{
					SPClass: "death_knight",
					SPSpec:  "frostDk",
				},
			},
			Name: "Razoon",
			URL:  "https://www.twitch.tv/razoonxz",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCSNjHLil0g8B4m2nlV_Pzfw",
				},
			},
		},
		{
			Slug:      "dariuswotlk",
			MainClass: "death_knight",
			ClassList: []string{
				"death_knight",
			},
			SpecList: []specList{
				{
					SPClass: "death_knight",
					SPSpec:  "unholy",
				},
			},
			Name: "Darius",
			URL:  "https://www.twitch.tv/dariuswotlk",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCNOmVPwbJH0XCCN7B1Sq8sg",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/MfRdDJQCAv",
				},
			},
		},
		{
			Slug:      "gargoylemvp",
			MainClass: "death_knight",
			ClassList: []string{
				"death_knight",
			},
			SpecList: []specList{
				{
					SPClass: "death_knight",
					SPSpec:  "unholy",
				},
			},
			Name:  "Gargooyle",
			URL:   "https://www.twitch.tv/gargoylemvp",
			Links: []socialLink{},
		},
		{
			Slug:      "cerva",
			MainClass: "death_knight",
			ClassList: []string{
				"death_knight",
			},
			Name: "Cerva",
			URL:  "https://www.twitch.tv/cerva",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCI2k32-vTf20BJHY6qnTB7w?view_as=subscriber",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/cervatv/",
				},
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Cervatv",
				},
			},
		},
		{
			Slug:      "losiro",
			MainClass: "death_knight",
			ClassList: []string{
				"death_knight",
			},
			Name: "Losiro",
			URL:  "https://www.twitch.tv/losiro",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/losirowow",
				},
			},
		},
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
			SpecList: []specList{
				{
					SPClass: "death_knight",
					SPSpec:  "unholy",
				},
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
			SpecList: []specList{
				{
					SPClass: "death_knight",
					SPSpec:  "unholy",
				},
				{
					SPClass: "warrior",
					SPSpec:  "arms",
				},
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
			Slug:      "creationlawlz",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "feral",
				},
			},
			Name:  "Creationlawlz",
			URL:   "https://www.twitch.tv/creationlawlz",
			Links: []socialLink{},
		},
		{
			Slug:      "xnatres",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "balance",
				},
			},
			Name:  "Xnatres",
			URL:   "https://www.twitch.tv/xnatres",
			Links: []socialLink{},
		},
		{
			Slug:      "simbo_feral",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "feral",
				},
			},
			Name:  "Simboferal",
			URL:   "https://www.twitch.tv/simbo_feral",
			Links: []socialLink{},
		},
		{
			Slug:      "bean",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "feral",
				},
			},
			Name: "Bean",
			URL:  "https://www.twitch.tv/bean",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/beantwitch",
				},
				{
					Slug: "youtube",
					URL:  "https://youtube.com/beantwitch",
				},
			},
		},
		{
			Slug:      "shyxy_tv",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "feral",
				},
			},
			Name: "Shyxy",
			URL:  "https://www.twitch.tv/shyxy_tv",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/ShyxyTV",
				},
			},
		},
		{
			Slug:      "moonfirebeam",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "balance",
				},
			},
			Name: "Moonfirebeam",
			URL:  "https://www.twitch.tv/moonfirebeam",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/moonfirebeam",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCsSYn11a0R4wyZUHyd57yIA",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/Z8TgJy6USv",
				},
			},
		},
		{
			Slug:      "spottman",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "feral",
				},
			},
			Name: "Spottman",
			URL:  "https://www.twitch.tv/spottman",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/SpottmanTV",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCQCZMWxDuofOStBWprVlK1Q",
				},
			},
		},
		{
			Slug:      "asgarathpvp",
			MainClass: "druid",
			ClassList: []string{
				"druid",
			},
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "restorationDruid",
				},
			},
			Name: "Asgarath",
			URL:  "https://www.twitch.tv/asgarathpvp",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Asgarathpvp",
				},
			},
		},
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
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "feral",
				},
			},
			Name: "Tonyferal",
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
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "restorationDruid",
				},
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
				{
					SPClass: "priest",
					SPSpec:  "holyPriest",
				},
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
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "feral",
				},
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
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "feral",
				},
				{
					SPClass: "shaman",
					SPSpec:  "elemental",
				},
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
			SpecList: []specList{
				{
					SPClass: "warrior",
					SPSpec:  "arms",
				},
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
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "restorationDruid",
				},
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
			SpecList: []specList{
				{
					SPClass: "druid",
					SPSpec:  "restorationDruid",
				},
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
		//EVOKER
		{
			Slug:      "bankstair",
			MainClass: "evoker",
			ClassList: []string{
				"evoker",
				"mage",
				"paladin",
			},
			Name: "Bankstair",
			URL:  "https://www.twitch.tv/bankstair",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCwp2cj0mC7ws4PuF4AssweA",
				},
			},
		},
		//HUNTER
		{
			Slug:      "namesix1",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			SpecList: []specList{
				{
					SPClass: "hunter",
					SPSpec:  "mm",
				},
			},
			Name: "Nayressa",
			URL:  "https://www.twitch.tv/namesix1",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/tibort5",
				},
			},
		},
		{
			Slug:      "homerjay_tv",
			MainClass: "hunter",
			ClassList: []string{
				"hunter",
			},
			SpecList: []specList{
				{
					SPClass: "hunter",
					SPSpec:  "mm",
				},
			},
			Name: "Homerjay",
			URL:  "https://www.twitch.tv/homerjay_tv",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCI8RZA6ZYtC_c0syNVUwqYg",
				},
			},
		},
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
			SpecList: []specList{
				{
					SPClass: "hunter",
					SPSpec:  "survival",
				},
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
			SpecList: []specList{
				{
					SPClass: "hunter",
					SPSpec:  "mm",
				},
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
			Slug:      "moopz",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			SpecList: []specList{
				{
					SPClass: "mage",
					SPSpec:  "frostMage",
				},
			},
			Name:  "Moopz",
			URL:   "https://www.twitch.tv/moopz",
			Links: []socialLink{},
		},
		{
			Slug:      "rivahlol",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			SpecList: []specList{
				{
					SPClass: "mage",
					SPSpec:  "frostMage",
				},
			},
			Name: "Rivah",
			URL:  "https://www.twitch.tv/rivahlol",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/RivahTBC",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCF7oCMiG304KSR1BbkutIzg?view_as=subscriber",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/KfTrWvBur2",
				},
			},
		},
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
			Name:  "Wealthyman",
			URL:   "https://www.twitch.tv/wealthymanx",
			Links: []socialLink{},
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
			},
		},
		{
			Slug:      "hansol",
			MainClass: "mage",
			ClassList: []string{
				"mage",
			},
			SpecList: []specList{
				{
					SPClass: "mage",
					SPSpec:  "fire",
				},
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
			SpecList: []specList{
				{
					SPClass: "mage",
					SPSpec:  "frostMage",
				},
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
			SpecList: []specList{
				{
					SPClass: "mage",
					SPSpec:  "fire",
				},
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
			SpecList: []specList{
				{
					SPClass: "mage",
					SPSpec:  "frostMage",
				},
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
			SpecList: []specList{
				{
					SPClass: "mage",
					SPSpec:  "frostMage",
				},
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
					URL:  "https://twitter.com/Xaryu",
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
			Slug:      "mysticallx",
			MainClass: "monk",
			ClassList: []string{
				"monk",
			},
			SpecList: []specList{
				{
					SPClass: "monk",
					SPSpec:  "mistweaver",
				},
			},
			Name: "Mysticallx",
			URL:  "https://www.twitch.tv/mysticallx",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/MysticalTheMonk",
				},
			},
		},
		{
			Slug:      "panzerhenk",
			MainClass: "monk",
			ClassList: []string{
				"monk",
			},
			SpecList: []specList{
				{
					SPClass: "monk",
					SPSpec:  "mistweaver",
				},
			},
			Name: "Panzerhenk",
			URL:  "https://www.twitch.tv/panzerhenk",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/panzerhenk",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/panzerhenk/",
				},
			},
		},
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
			SpecList: []specList{
				{
					SPClass: "monk",
					SPSpec:  "mistweaver",
				},
				{
					SPClass: "shaman",
					SPSpec:  "restorationShaman",
				},
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
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
				"rogue",
				"monk",
				"evoker",
				"priest",
				"warlock",
				"hunter",
				"death_knight",
				"demon_hunter",
				"mage",
				"paladin",
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
			Slug:      "h0lyundead",
			MainClass: "priest",
			ClassList: []string{
				"priest",
				"rogue",
			},
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "disc",
				},
			},
			Name: "H0lyundead",
			URL:  "https://www.twitch.tv/h0lyundead",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCIxXLM7vT8JKc-UF9cJ9RTw?view_as=subscriber",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/BxgGkZ7nEE",
				},
			},
		},
		{
			Slug:      "earpugs",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "disc",
				},
			},
			Name: "Earpugs",
			URL:  "https://www.twitch.tv/earpugs",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCTxM4AZCZKfyxNH8l2hm4uQ",
				},
			},
		},
		{
			Slug:      "tempestxgg",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "disc",
				},
			},
			Name:  "Tempest",
			URL:   "https://www.twitch.tv/tempestxgg",
			Links: []socialLink{},
		},
		{
			Slug:      "ascalontv",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "shadow",
				},
			},
			Name: "Ascalontv",
			URL:  "https://www.twitch.tv/ascalontv",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/AscalonTV",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC6FKFoJWGtZoKIDZRSXg78A",
				},
				{
					Slug: "discord",
					URL:  "https://discord.com/invite/uepKAsY",
				},
			},
		},
		{
			Slug:      "hozitojones",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "disc",
				},
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
			Slug:      "aphobiagaming",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "shadow",
				},
			},
			Name: "Aphobia",
			URL:  "https://www.twitch.tv/aphobiagaming",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/eddphobia",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/eddphobia91/",
				},
			},
		},
		{
			Slug:      "olliektv",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "disc",
				},
			},
			Name: "Olliektv",
			URL:  "https://www.twitch.tv/olliektv",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/GbQrSUyR",
				},
			},
		},
		{
			Slug:      "stahpsp",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "shadow",
				},
			},
			Name: "Stahpsp",
			URL:  "https://www.twitch.tv/stahpsp",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/stahpsp",
				},
			},
		},
		{
			Slug:      "krawnzlol",
			MainClass: "priest",
			ClassList: []string{
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "disc",
				},
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
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "disc",
				},
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
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "shadow",
				},
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
			SpecList: []specList{
				{
					SPClass: "priest",
					SPSpec:  "disc",
				},
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
			Slug:      "chokopapatv",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
			},
			Name:  "Chokopapatv",
			URL:   "https://www.twitch.tv/chokopapatv",
			Links: []socialLink{},
		},
		{
			Slug:      "preghierax",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "ret",
				},
			},
			Name:  "Preghierax",
			URL:   "https://www.twitch.tv/preghierax",
			Links: []socialLink{},
		},
		{
			Slug:      "holypalaswe1",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
			},
			Name: "Holypalaswe1",
			URL:  "https://www.twitch.tv/holypalaswe1",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/geVNEUue7C",
				},
				{
					Slug: "youtube",
					URL:  "https://youtube.com/user/holypalaswe/videos",
				},
			},
		},
		{
			Slug:      "flowstateswow",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
			},
			Name:  "Flowstateswow",
			URL:   "https://www.twitch.tv/flowstateswow",
			Links: []socialLink{},
		},
		{
			Slug:      "pizzahunter2009",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "ret",
				},
			},
			Name: "Holysaxyboy",
			URL:  "https://www.twitch.tv/pizzahunter2009",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/W8M5NMF",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UC1yMtI5eIE4cTex7KTfxb0w",
				},
			},
		},
		{
			Slug:      "waynxt",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
			},
			Name: "Wayne",
			URL:  "https://www.twitch.tv/waynxt",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/user/BarneyGamingCZ/videos",
				},
			},
		},
		{
			Slug:      "flyntv_",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
			},
			Name: "Saori",
			URL:  "https://www.twitch.tv/flyntv_",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/SSJzzzupBQ",
				},
			},
		},
		{
			Slug:      "soumi",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
			},
			Name: "Soumi",
			URL:  "https://www.twitch.tv/soumi",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/3882XAKwyg",
				},
			},
		},
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
			Slug:      "crusader3455",
			MainClass: "paladin",
			ClassList: []string{
				"paladin",
			},
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "ret",
				},
			},
			Name: "Crusader",
			URL:  "https://www.twitch.tv/crusader3455",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/itsmikeowens",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/itsmikeowens/",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCzoNEIaIy1Q-VrrU1P-5vpw",
				},
			},
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
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "ret",
				},
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
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
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
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "ret",
				},
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
			SpecList: []specList{
				{
					SPClass: "paladin",
					SPSpec:  "ret",
				},
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
			Slug:      "reqtbc",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			SpecList: []specList{
				{
					SPClass: "rogue",
					SPSpec:  "sub",
				},
			},
			Name: "Reqtbc",
			URL:  "https://www.twitch.tv/reqtbc",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/DVQYXk4",
				},
			},
		},
		{
			Slug:      "perplexity",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			SpecList: []specList{
				{
					SPClass: "rogue",
					SPSpec:  "sub",
				},
			},
			Name: "Perplexity",
			URL:  "https://www.twitch.tv/perplexity",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/aXryrUK",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/watch?v=fLkZuz7kjyU&ab_channel=perplexity",
				},
			},
		},
		{
			Slug:      "avizura",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			SpecList: []specList{
				{
					SPClass: "rogue",
					SPSpec:  "sub",
				},
			},
			Name: "Avizura",
			URL:  "https://www.twitch.tv/avizura",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/avizuray",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/Avizura/featured?view_as=subscriber",
				},
			},
		},
		{
			Slug:      "petraxs",
			MainClass: "rogue",
			ClassList: []string{
				"rogue",
			},
			SpecList: []specList{
				{
					SPClass: "rogue",
					SPSpec:  "sub",
				},
			},
			Name: "Petraxs",
			URL:  "https://www.twitch.tv/petraxs",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/c/Petraxs",
				},
			},
		},
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
			SpecList: []specList{
				{
					SPClass: "rogue",
					SPSpec:  "sub",
				},
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
			SpecList: []specList{
				{
					SPClass: "rogue",
					SPSpec:  "sub",
				},
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
			SpecList: []specList{
				{
					SPClass: "rogue",
					SPSpec:  "assa",
				},
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
			Slug:      "groblinwow",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "restorationShaman",
				},
			},
			Name: "Groblin",
			URL:  "https://twitch.tv/groblinwow",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/SFJ8eNnDYS",
				},
			},
		},
		{
			Slug:      "zakiwak1",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "restorationShaman",
				},
			},
			Name: "Zakiwaki",
			URL:  "https://twitch.tv/zakiwak1",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/8akXAsnQ4p",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCaXjkbB_eCUpQo76raxYa1A",
				},
			},
		},
		{
			Slug:      "traktorch0",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "enhancement",
				},
			},
			Name:  "Traktorcho",
			URL:   "https://www.twitch.tv/traktorch0",
			Links: []socialLink{},
		},
		{
			Slug:      "skilerxtv",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "elemental",
				},
			},
			Name: "Skiler",
			URL:  "https://www.twitch.tv/skilerxtv",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/yewtFpej6j",
				},
			},
		},
		{
			Slug:      "blackbettytv",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "elemental",
				},
			},
			Name:  "Blackbetty",
			URL:   "https://www.twitch.tv/blackbettytv",
			Links: []socialLink{},
		},
		{
			Slug:      "novoz",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "elemental",
				},
			},
			Name:  "Novoz",
			URL:   "https://www.twitch.tv/novoz",
			Links: []socialLink{},
		},
		{
			Slug:      "hrkzz",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "elemental",
				},
			},
			Name: "Hrkzz",
			URL:  "https://www.twitch.tv/hrkzz",
			Links: []socialLink{
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/rn.fitt/?hl=fr",
				},
			},
		},
		{
			Slug:      "warbla",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "restorationShaman",
				},
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
			Name:  "Kolowavex",
			URL:   "https://www.twitch.tv/kolowavex",
			Links: []socialLink{},
		},
		{
			Slug:      "tiqqlethis",
			MainClass: "shaman",
			ClassList: []string{
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "enhancement",
				},
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
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "enhancement",
				},
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
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "restorationShaman",
				},
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
			MainClass: "paladin",
			ClassList: []string{
				"shaman",
				"monk",
				"druid",
				"priest",
				"paladin",
				"evoker",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "restorationShaman",
				},
				{
					SPClass: "druid",
					SPSpec:  "restorationDruid",
				},
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
				{
					SPClass: "priest",
					SPSpec:  "holyPriest",
				},
				{
					SPClass: "monk",
					SPSpec:  "mistweaver",
				},
				{
					SPClass: "evoker",
					SPSpec:  "preservation",
				},
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
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "elemental",
				},
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
				"paladin",
				"monk",
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "monk",
					SPSpec:  "mistweaver",
				},
				{
					SPClass: "shaman",
					SPSpec:  "restorationShaman",
				},
				{
					SPClass: "druid",
					SPSpec:  "restorationDruid",
				},
				{
					SPClass: "paladin",
					SPSpec:  "holyPaladin",
				},
				{
					SPClass: "priest",
					SPSpec:  "holyPriest",
				},
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
			Slug:      "infernion",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
				"warrior",
				"priest",
				"mage",
			},
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "affliction",
				},
			},
			Name: "infernion",
			URL:  "https://www.twitch.tv/infernion",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/Infernionwow",
				},
				{
					Slug: "discord",
					URL:  "https://discord.gg/R8pKuhhba3",
				},
				{
					Slug: "instagram",
					URL:  "https://www.instagram.com/jespergjessen/",
				},
			},
		},
		{
			Slug:      "chanx",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "destruction",
				},
			},
			Name:  "chanx",
			URL:   "https://www.twitch.tv/chanx",
			Links: []socialLink{},
		},
		{
			Slug:      "kerosineyo",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "destruction",
				},
			},
			Name:  "Kerosineyo",
			URL:   "https://www.twitch.tv/kerosineyo",
			Links: []socialLink{},
		},
		{
			Slug:      "scryna",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "destruction",
				},
			},
			Name: "Scryna",
			URL:  "https://www.twitch.tv/scryna",
			Links: []socialLink{
				{
					Slug: "twitter",
					URL:  "https://twitter.com/scryna",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/@scryna",
				},
			},
		},
		{
			Slug:      "reverencewarlock",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "destruction",
				},
			},
			Name:  "Reverence",
			URL:   "https://www.twitch.tv/reverencewarlock",
			Links: []socialLink{},
		},
		{
			Slug:      "bualock",
			MainClass: "warlock",
			ClassList: []string{
				"warlock",
			},
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "destruction",
				},
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
				"priest",
			},
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "destruction",
				},
				{
					SPClass: "priest",
					SPSpec:  "shadow",
				},
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
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "destruction",
				},
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
				// "priest",
			},
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "destruction",
				},
				// {
				// 	SPClass: "priest",
				// 	SPSpec:  "shadow",
				// },
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
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "affliction",
				},
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
			SpecList: []specList{
				{
					SPClass: "warlock",
					SPSpec:  "destruction",
				},
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
				"shaman",
			},
			SpecList: []specList{
				{
					SPClass: "shaman",
					SPSpec:  "elemental",
				},
				{
					SPClass: "warlock",
					SPSpec:  "affliction",
				},
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
			Slug:      "kostickatv",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			SpecList: []specList{
				{
					SPClass: "warrior",
					SPSpec:  "protection",
				},
			},
			Name: "Kosticka",
			URL:  "https://www.twitch.tv/kostickatv",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/ejWtuhR",
				},
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCT0QT0ebAi1U0myAwuoowvQ",
				},
			},
		},
		{
			Slug:      "bcxwarr",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			SpecList: []specList{
				{
					SPClass: "warrior",
					SPSpec:  "arms",
				},
			},
			Name: "Bcw",
			URL:  "https://www.twitch.tv/bcxwarr",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/PwCBmNeGDF",
				},
			},
		},
		{
			Slug:      "torstenstock",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			SpecList: []specList{
				{
					SPClass: "warrior",
					SPSpec:  "arms",
				},
			},
			Name: "Torstenstock",
			URL:  "https://www.twitch.tv/torstenstock",
			Links: []socialLink{
				{
					Slug: "discord",
					URL:  "https://discord.gg/jDVkDEZA",
				},
			},
		},
		{
			Slug:      "pvelordtv",
			MainClass: "warrior",
			ClassList: []string{
				"warrior",
			},
			Name: "Pvelordtv",
			URL:  "https://www.twitch.tv/pvelordtv",
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtube.com/channel/UCUQMAH2NXCM1GKznX5QYlMQ",
				},
			},
		},
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
			SpecList: []specList{
				{
					SPClass: "warrior",
					SPSpec:  "arms",
				},
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
			SpecList: []specList{
				{
					SPClass: "warrior",
					SPSpec:  "arms",
				},
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
			SpecList: []specList{
				{
					SPClass: "warrior",
					SPSpec:  "arms",
				},
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

	// sort.Slice(streams, func(a, b int) bool {
	// 	return strings.Compare(streams[a].Name, streams[b].Name) < 0
	// })

	// for _, r := range streams {
	// 	sort.Slice(r.Links, func(a, b int) bool {
	// 		return strings.Compare(r.Links[a].Slug, r.Links[b].Slug) < 0
	// 	})
	// }

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
