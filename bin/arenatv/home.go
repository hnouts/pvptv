package main

import (
	"fmt"

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
	ctx.Page().SetDescription("all of wow pvp streaming in one place")
}

func (p *homePage) Render() app.UI {
	return newPage().
		Title("PvPtv.io").
		Icon("/web/logo.png").
		Index(
			newIndexLink().Title("Welcome!").Href("#welcome-to-pvptv"),
			newIndexLink().Title("How to find a streamer that plays my class?").Href("#help"),
			newIndexLink().Title("Why is my twitch screen purple?").Href("#why-is-my-twitch-screen-purple"),
			newIndexLink().Title("More resources").Href("#more-resources"),
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
							Text("PvPtv.io is an app that was created in order to help the world of warcraft pvp community find the best arena streamers at any moment. It provides a list of the best known wow arena's streamers for each class, and casts their stream directly from twitch.com"),
						app.Br(),
						app.P().Class("text").
							Text("The idea behind this web application was to finally have an easy solution to the eternal question posted on pvp communities:"),
						app.P().Class("text").
							Text("\"Do you know a good streamer for x class?\" Now you do! And among the vast list of streamers stored on PvPtv.io, there will always be a gladiator streaming! (hopefully)"),
						app.P().
							Class("text").
							Text("You can contribute to the list of streamers by filling the following form:"),
						app.A().Class("center-link").Href("https://forms.gle/y5wj532gvtgwWkca6").Text("-> suggest a streamer <-"),
					),
					app.Div().ID("help").Body(
						app.H2().
							Class("title").
							Text("How to find a streamer that plays my class?"),
						app.P().
							Class("text").
							Text("It's very easy! Select a class and the app will list streamers known to play this class. Sometimes streamers will be playing a different class... But the app is trying its best to determine what the streamer is playing right now!"),
						app.Br(),
						app.P().
							Class("text").
							Text("To help you choose a stream, icons are displayed next to the streamer's name."),
						app.Div().
							Class("help-icon").
							Body(
								app.Raw(fmt.Sprintf(subRogueSVG, 25, 25)),
								app.P().
									Text("Known to main this given class and spec"),
							),
						app.Div().
							Class("help-icon alt-border").
							Body(
								app.Raw(fmt.Sprintf(subRogueSVG, 25, 25)),
								app.P().
									Text("Known to play this given class and spec from time to time"),
							),

						app.Div().
							Class("help-icon border").
							Body(
								app.Raw(fmt.Sprintf(websiteSVG, 25, 25)),
								app.P().
									Text("Could be playing anything"),
							),
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
							Text("Twitch's purple screen error occurs when you're watching live streams from a website other than twitch.com. It can be triggered by your ad-blocker and can be avoided if you disable it on PvPtv.io. Do not worry! PvPtv.io has zero ads!"),
					),
					app.Div().ID("more-resources").Body(
						app.H2().
							Class("title").
							Text("Pvp Community resources"),
						app.P().
							Class("text").
							Text("Links to other world of warcraft pvp resources."),
						app.A().Class().Href("https://www.reddit.com/r/worldofpvp").Text("r/worldofpvp"),
						app.Br(),
						app.A().Class().Href("https://www.pvpleaderboard.com/").Text("pvpleaderboard"),
						app.Br(),
						app.A().Class().Href("https://pvp.subcreation.net/").Text("pvp.subcreation"),
						app.Br(),
						app.A().Class().Href("https://check-pvp.fr/").Text("checkpvp"),
						app.Br(),
					),
					app.Div().ID("special-thanks").Body(
						app.H2().
							Class("title").
							Text("Special thanks"),
						app.P().
							Class("text").
							Text("PvPtv.io is built with the amazing golang framework go-app-dev! You can contact the author on twitter or visit the frameworks documentation."),
						app.A().Class().Href("https://twitter.com/jonhymaxoo").Text("Maxence Twitter"),
						app.Br(),
						app.A().Class().Href("https://go-app.dev/").Text("Go-App-Dev Documentation"),
						app.Br(),
						app.P().
							Class("text").
							Text("Thanks also to Dmachine and his rich list of arena streamers that was a very useful source to me and an inspiration to this project."),
						app.A().Class().Href("https://twitter.com/Dmachine_").Text("DMachine Twitter"),
						app.Br(),
						app.A().Class().Href("https://www.reddit.com/r/worldofpvp/comments/e9uukx/list_of_wow_pvp_streamers_organized_by_classspec/").Text("Reddit list of streamers"),
					),
				),
		)
}
