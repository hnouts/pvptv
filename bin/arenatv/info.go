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
	i.currentCard = -1

	ticker := time.NewTicker(cardHiddenDuration)
	ctx.Async(func() {
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				ctx.Dispatch(func(ctx app.Context) {
					if i.isCardVisible {
						ticker.Reset(cardHiddenDuration)
						i.isCardVisible = false
					} else {
						ticker.Reset(cardVisibleDuration)
						i.isCardVisible = true
						i.showNewCard(ctx)
					}
				})
			}
		}
	})
}

func (i *info) OnUpdate(ctx app.Context) {
}

func (i *info) showNewCard(ctx app.Context) {
	count := len(i.Istream.Cards)
	if count == 0 {
		i.currentCard = -1
		return
	}

	i.currentCard++
	if i.currentCard >= count {
		i.currentCard = 0
	}
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
					app.Div().Class("info-title-separator"),
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
			app.Range(i.Istream.Cards).Slice(func(j int) app.UI {
				cardVisibility := ""
				if j == i.currentCard && i.Iplaying && i.isCardVisible {
					cardVisibility = "info-card-show"
				}

				return app.P().
					Class("info-card").
					Class("glow").
					Class("focus").
					Class(cardVisibility).
					Text(i.Istream.Cards[j])
			}),
		)
}
