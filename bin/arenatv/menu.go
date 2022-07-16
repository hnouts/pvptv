package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type menu struct {
	app.Compo

	Iclass string

	appInstallable bool
}

func newMenu() *menu {
	return &menu{}
}

func (m *menu) Class(v string) *menu {
	m.Iclass = app.AppendClass(m.Iclass, v)
	return m
}

func (m *menu) OnNav(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) OnAppInstallChange(ctx app.Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) Render() app.UI {
	linkClass := "link heading fit unselectable"

	return app.Div().
		Class("nav").
		Class("fill").
		Class("unselectable").
		Class(m.Iclass).
		Body(
			ui.Stack().
				Class("app-title").
				Class("hspace-out").
				Middle().
				Content(
					app.Header().
						Body(
							app.A().
								Class("hApp").
								Class("focus").
								Class("glow").
								Href("/").
								Text("Arenatv"),
						),
				),
			app.Nav().Class("nav-content").
				Body(
					app.Div().
						Class("nav-streams").
						Body(
							ui.Stack().
								Class("nav-streams-stack").
								Middle().
								Content(
									app.Div().
										Class("hspace-out").
										Body(
											ui.Link().
												Class(linkClass).
												Label("Priest").
												Href("/priest"),
											ui.Link().
												Class(linkClass).
												Label("Paladin").
												Href("/paladin"),
											ui.Link().
												Class(linkClass).
												Label("Rogue").
												Href("/rogue"),
										),
								),
						),
					// app.Div().
					// 	Class("nav-streams").
					// 	Body(
					// 		ui.Link().
					// 			Class(linkClass).
					// 			Label("Priest").
					// 			Href("/priest"),
					// 		ui.Link().
					// 			Class(linkClass).
					// 			Label("Paladin").
					// 			Href("/paladin"),
					// 		ui.Link().
					// 			Class(linkClass).
					// 			Label("Rogue").
					// 			Href("/rogue")),

					// app.Div().Class("info-title-separator"),

					app.Div().
						Class("nav-support").
						Class("hspace-out").
						Body(
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(coffeeSVG)).
								Label("Buy me a coffee").
								Href(buyMeACoffeeURL),
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(githubSVG)).
								Label("GitHub").
								Href(githubURL),
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(twitterSVG)).
								Label("Twitter").
								Href(twitterURL),
						),
				),
		)
}
