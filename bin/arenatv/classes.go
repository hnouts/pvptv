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
			Slug:    "priest",
			Name:    "Priest",
			Classic: true,
		},
		{
			Slug:    "paladin",
			Name:    "Paladin",
			Classic: true,
		},
		{
			Slug:    "rogue",
			Name:    "Rogue",
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
