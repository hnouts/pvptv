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
