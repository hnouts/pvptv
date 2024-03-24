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
	Svg            string
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
			Slug: "death_knight",
			Name: "Death Knight",
			Svg:  dkSVG,
		},
		{
			Slug: "demon_hunter",
			Name: "Demon Hunter",
			Svg:  dhSVG,
		},
		{
			Slug: "druid",
			Name: "Druid",
			Svg:  druidSVG,
		},
		{
			Slug: "evoker",
			Name: "Evoker",
			Svg:  evokerSVG,
		},
		{
			Slug: "hunter",
			Name: "Hunter",
			Svg:  hunterSVG,
		},
		{
			Slug: "mage",
			Name: "Mage",
			Svg:  mageSVG,
		},
		{
			Slug: "monk",
			Name: "Monk",
			Svg:  monkSVG,
		},
		{
			Slug: "paladin",
			Name: "Paladin",
			Svg:  paladinSVG,
		},
		{
			Slug: "priest",
			Name: "Priest",
			Svg:  priestSVG,
		},
		{
			Slug: "rogue",
			Name: "Rogue",
			Svg:  rogueSVG,
		},
		{
			Slug: "shaman",
			Name: "Shaman",
			Svg:  shamanSVG,
		},
		{
			Slug: "warlock",
			Name: "Warlock",
			Svg:  warlockSVG,
		},
		{
			Slug: "warrior",
			Name: "Warrior",
			Svg:  warriorSVG,
		},
		{
			Slug: "plunderstorm",
			Name: "Plunderstorm",
			Svg:  plunderstormSVG,
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
}
