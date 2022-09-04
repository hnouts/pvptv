package main

import (
	// "github.com/maxence-charriere/go-app/v9/pkg/analytics"
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
	ctx.Page().SetTitle("PvPtv.io")
	ctx.Page().SetDescription("all wow pvp streamer in one place")
	// analytics.Page("home", nil)
}

func (p *homePage) Render() app.UI {
	return newPage().
		Title("pvptv.io").
		Icon("/web/logo.png").
		Index(
			newIndexLink().Title("Welcome!").Href("#welcome-to-pvptv"),
			newIndexLink().Title("Why is my twitch screen purple?").Href("#why-is-my-twitch-screen-purple"),
			newIndexLink().Title("Special thanks"),
		).
		Content(
			ui.Flow().
				StretchItems().
				Spacing(84).
				Content(
					app.Div().ID("welcome-to-pvptv").Body(
						app.H2().
							Class("title").
							Text("Welcome!"),
						app.P().
							Class("text").
							Text("Pvptv.io is an app that was created in order to help the world of warcraft pvp community find the best arena streamers at any moment. It provides a list of the best known wow arena's streamers for each class, and casts their stream directly from twitch.com."),
						app.Br(),
						app.P().Class("text").
							Text("The motivation behind this web application was to finally have an easy solution to the eternal question posted on pvp communities:"),
						app.Div().Class("questions").Body(
							app.Img().
							Class("welcome-example").
							Alt("eternal-question-1").
							Src("/web/lfstream1.png"),
						app.Img().
							Class("welcome-example").
							Alt("eternal-question-2").
							Src("/web/lfstream1.png"),
						),
						app.P().Class("text").
							Text("\"Do you know a good streamer for x class?\" Now you do! And among the vast list of streamers stored on Pvptv.io, there will always be a gladiator streaming! (hopefully)"),
					),
					app.Div().ID("why-is-my-twitch-screen-purple").Body(
						app.H2().
							Class("title").
							Text("Why is my twitch screen purple?"),
						app.Div().Class().Body(
							app.Img().
							Class("twitch-pod").
							Alt("twitch-purple-screen").
							Src("/web/twitch-purple-screen.png"),
						),
						app.P().
							Class("text").
							Text("Twitch's purple screen error might occurs when you're watching live streams from another website than twitch.com. It can be triggered by your ad-blockers and can be avoided if you disable it on Pvptv.io. Do not worry! Pvptv.io has zero ad!"),
					),
					app.Div().ID("special-thanks").Body(
						app.H2().
							Class("title").
							Text("Special thanks"),
						app.P().
							Class("text").
							Text("Pvptv.io is built with the amazing golang framework go-app-dev! You can contact the author on twitter or visit the frameworks documentation"),
						app.A().Class().Href("https://twitter.com/jonhymaxoo").Text("Maxence Twitter"),
						app.Br(),
						app.A().Class().Href("https://go-app.dev/").Text("Go-App-Dev Documentation"),
					),
				),
		)
}
