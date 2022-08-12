package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

const (
	infoLinkIconSize    = 18
	cardVisibleDuration = time.Second * 10
	cardHiddenDuration  = time.Second * 5
)

type info struct {
	app.Compo

	Iclass   string
	Istream  liveStream
	Iplaying bool

	currentCard   int
	isCardVisible bool
}

func newInfo() *info {
	return &info{}
}

func (i *info) Class(v string) *info {
	if v == "" {
		return i
	}
	if i.Iclass != "" {
		i.Iclass += " "
	}
	i.Iclass += v
	return i
}

func (i *info) Stream(v liveStream) *info {
	i.Istream = v
	return i
}

func (i *info) Playing(v bool) *info {
	i.Iplaying = v
	return i
}

func (i *info) OnMount(ctx app.Context) {
}

func (i *info) OnUpdate(ctx app.Context) {
}

func (i *info) Render() app.UI {
	titleVisibility := "info-title-show"
	// if i.Iplaying {
	// 	titleVisibility = "info-title-show"
	// }

	// fmt.Println("CALLED Render Info :", i)

	return app.Article().
		Class("info").
		Class("fill").
		Class("no-overflow").
		Body(
			app.Header().
				Class("info-title").
				Class(titleVisibility).
				Class("center").
				Class("fit").
				Body(
					app.H1().
						Class("h1").
						Class("glow").
						Text(i.Istream.Name),
					app.If(len(i.Istream.Name) > 0,
						app.Div().Class("info-title-separator"),
					),
					ui.Stack().
						Class("info-links").
						Center().
						Middle().
						Content(
							app.Range(i.Istream.Links).Slice(func(j int) app.UI {
								l := i.Istream.Links[j]
								return newInfoLink().
									Help(fmt.Sprintf("Visit %s's %s.",
										strings.Title(i.Istream.Name),
										strings.Title(l.Slug),
									)).
									Href(l.URL).
									Icon(newSVGIcon().
										Size(infoLinkIconSize).
										RawSVG(socialIcon(l.Slug)))
							}),
						),
				),
		)
}
