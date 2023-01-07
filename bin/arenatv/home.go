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
						app.Div().Class("").Body(
							app.Img().
								Class("welcome-example").
								Alt("eternal-question-1").
								Src("/web/streams-question.png"),
							app.Img().
								Class("welcome-example").
								Alt("eternal-question-1").
								Src("/web/streams-question2.png"),
						),
						app.Div().Class("").Body(
							app.Img().
								Class("welcome-example").
								Alt("eternal-question-2").
								Src("/web/streams-question3.png"),
							app.Img().
								Class("welcome-example").
								Alt("eternal-question-2").
								Src("/web/streams-question4.png"),
						),
						app.P().Class("text").
							Text("\"Do you know a good streamer for x class?\" Now you do! And among the vast list of streamers stored on PvPtv.io, there will always be a gladiator streaming! (hopefully)"),
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

						// app.Div().
						// 	Class("help-icon border").
						// 	Body(
						// 		app.Raw(fmt.Sprintf(subRogueSVG, 38, 25)),
						// 		app.P().
						// 			Text("This streamer could be playing anything, but more likely to be playing this spec based on the current meta"),
						// 	),
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
							Text("Thanks also to reddit user Dmachine_Blizz and his rich list of arena streamers that was a very useful source to me and an inspiration to this project."),
						app.A().Class().Href("https://twitter.com/Dmachine_").Text("DMachine Twitter"),
						app.Br(),
						app.A().Class().Href("https://www.reddit.com/r/worldofpvp/comments/e9uukx/list_of_wow_pvp_streamers_organized_by_classspec/").Text("Reddit list of streamers"),
					),
				),
		)
}
