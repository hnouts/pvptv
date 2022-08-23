package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type menu struct {
	app.Compo

	Iclass string

	appInstallable  bool
	updateAvailable bool
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
	linkClass := "link heading fit icon-circle"

	return app.Div().
		Class("nav").
		Class("fill").
		// Class("unselectable").
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
								Text("Pvptv.io"),
						),
				),

			app.Div().
				Class("hspace-out").
				Body(
					app.If(m.appInstallable,
						ui.Link().
							Class(linkClass).
							Icon(downloadSVG).
							Label("Install").
							OnClick(m.installApp),
					),
					app.If(m.updateAvailable,
						ui.Link().
							Class(linkClass).
							Icon(downloadSVG).
							Label("Update").
							OnClick(m.updateApp),
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
												Label("Death Knight").
												Icon(dkSVG).
												Href("/death_knight"),
											ui.Link().
												Class(linkClass).
												Label("Demon Hunter").
												Icon(dhSVG).
												Href("/demon_hunter"),
											ui.Link().
												Class(linkClass).
												Label("Druid").
												Icon(druidSVG).
												Href("/druid"),
											ui.Link().
												Class(linkClass).
												Label("Hunter").
												Icon(hunterSVG).
												Href("/hunter"),
											ui.Link().
												Class(linkClass).
												Label("Mage").
												Icon(mageSVG).
												Href("/mage"),
											ui.Link().
												Class(linkClass).
												Label("Monk").
												Icon(monkSVG).
												Href("/monk"),
											ui.Link().
												Class(linkClass).
												Label("Paladin").
												Icon(paladinSVG).
												Href("/paladin"),
											ui.Link().
												Class(linkClass).
												Label("Priest").
												Icon(priestSVG).
												Href("/priest"),
											ui.Link().
												Class(linkClass).
												Label("Rogue").
												Icon(rogueSVG).
												Href("/rogue"),
											ui.Link().
												Class(linkClass).
												Label("Shaman").
												Icon(shamanSVG).
												Href("/shaman"),
											ui.Link().
												Class(linkClass).
												Label("Warlock").
												Icon(warlockSVG).
												Href("/warlock"),
											ui.Link().
												Class(linkClass).
												Label("Warrior").
												Icon(warriorSVG).
												Href("/warrior"),
										),
								),
						),
					app.Div().
						Class("nav-support-home").
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

func (m *menu) installApp(ctx app.Context, e app.Event) {
	ctx.NewAction(installApp)
}
func (m *menu) updateApp(ctx app.Context, e app.Event) {
	ctx.NewAction(updateApp)
}
