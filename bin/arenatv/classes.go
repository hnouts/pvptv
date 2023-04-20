package main

import (
	"sort"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type classes struct {
	app.Compo
	Name           string
	Slug           string
	Svg			   string
	lives          []liveStream
	current        liveStream
	Classic        bool
	Iclass         string
	IliveStreams   []liveStream
	IcurrentStream liveStream

	isFirstLoad bool
}

func newClasses() *classes {
	return &classes{}
}

func getAllClasses() []classes {
	classes := []classes{
		{
			Slug:    "death_knight",
			Name:    "Death Knight",
			Svg:	 dkSVG,
			Classic: true,
		},
		{
			Slug:    "demon_hunter",
			Name:    "Demon Hunter",
			Svg:	 dhSVG,
			Classic: false,
		},
		{
			Slug:    "druid",
			Name:    "Druid",
			Svg:	 druidSVG,
			Classic: true,
		},
		{
			Slug:    "evoker",
			Name:    "Evoker",
			Svg:	 evokerSVG,
			Classic: false,
		},
		{
			Slug:    "hunter",
			Name:    "Hunter",
			Svg:	 hunterSVG,
			Classic: true,
		},
		{
			Slug:    "mage",
			Name:    "Mage",
			Svg:	 mageSVG,
			Classic: true,
		},
		{
			Slug:    "monk",
			Name:    "Monk",
			Svg:	 monkSVG,
			Classic: false,
		},
		{
			Slug:    "paladin",
			Name:    "Paladin",
			Svg:	 paladinSVG,
			Classic: true,
		},
		{
			Slug:    "priest",
			Name:    "Priest",
			Svg:	 priestSVG,
			Classic: true,
		},
		{
			Slug:    "rogue",
			Name:    "Rogue",
			Svg:	 rogueSVG,
			Classic: true,
		},
		{
			Slug:    "shaman",
			Name:    "Shaman",
			Svg:	 shamanSVG,
			Classic: true,
		},
		{
			Slug:    "warlock",
			Name:    "Warlock",
			Svg:	 warlockSVG,
			Classic: true,
		},
		{
			Slug:    "warrior",
			Name:    "Warrior",
			Svg:	 warriorSVG,
			Classic: true,
		},
	}

	sort.Slice(classes, func(a, b int) bool {
		return strings.Compare(classes[a].Name, classes[b].Name) < 0
	})

	return classes
}

func (c *classes) Render() app.UI {
	return newPage().
		Title(c.Name)
	// Content(
	// 	ui.Shell().
	// 		Class("stream-shell").
	// 		Class("fill").
	// 		PaneWidth(menuWidth).
	// 		Menu(newNav().
	// 			CurrentClass(c.current.Class).
	// 			LiveStreams(c.lives).
	// 			CurrentStream(c.current)).
	// 		HamburgerMenu(newNav().
	// 			Class("background-overlay").
	// 			CurrentClass(c.current.Class).
	// 			LiveStreams(c.lives).
	// 			CurrentStream(c.current)),
	// )
}
