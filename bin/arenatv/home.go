package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type homePage struct {
	app.Compo
}

func newHomePage() *homePage {
	return &homePage{}
}

func (p *homePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Arenatv.io")
	ctx.Page().SetDescription("best website on azeroth")
	analytics.Page("home", nil)
}

func (p *homePage) Render() app.UI {
	return newPage().
		Index(
			newIndexLink().Title("What is arenatv?"),
			newIndexLink().Title("Twitch Purple screen of death"),
		).
		Content(
			ui.Flow().
				StretchItems().
				Spacing(84).
				Content(
					app.Div().Body(
						app.H1().
							Class("title").
							Text("Build a GUI with Go"),
						app.P().
							Class("text").
							Text("Just because Go and this package are really awesome!"),
					),
					app.Div().Body(
						app.H1().
							Class("title").
							Text("Build a GUI with Go"),
						app.P().
							Class("text").
							Text("Just because Go and this package are really awesome!"),
					),
					app.Div().Body(
						app.H1().
							Class("title").
							Text("Build a GUI with Go"),
						app.P().
							Class("text").
							Text("Just because Go and this package are really awesome!"),
					),
					app.Div().Body(
						app.H1().
							Class("title").
							Text("Build a GUI with Go"),
						app.P().
							Class("text").
							Text("Just because Go and this package are really awesome!"),
					),

					app.Div().Body(
						app.H1().
							Class("title").
							Text("Build a GUI with Go"),
						app.P().
							Class("text").
							Text("Just because Go and this package are really awesome!"),
					),
					app.Div().Body(
						app.H1().
							Class("title").
							Text("Build a GUI with Go"),
						app.P().
							Class("text").
							Text("Just because Go and this package are really awesome!"),
					),
					app.Div().Body(
						app.H1().
							Class("title").
							Text("Build a GUI with Go"),
						app.P().
							Class("text").
							Text("Just because Go and this package are really awesome!"),
					),
					app.Div().Body(
						app.H1().
							Class("title").
							Text("Build a GUI with Go"),
						app.P().
							Class("text").
							Text("Just because Go and this package are really awesome!"),
					),
				),
		)
}
