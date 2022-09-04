package main

import (
	"strconv"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type nav struct {
	app.Compo

	Iclass         string
	IcurrentClass  string
	IliveStreams   []liveStream
	IcurrentStream liveStream
	// classSvg       string
	isFirstLoad    bool
}

func newNav() *nav {
	return &nav{}
}

func (n *nav) Class(v string) *nav {
	if v == "" {
		return n
	}
	if n.Iclass != "" {
		n.Iclass += " "
	}
	n.Iclass += v
	return n
}

func (n *nav) CurrentClass(v string) *nav {
	n.IcurrentClass = v
	return n
}

func (n *nav) LiveStreams(v []liveStream) *nav {
	n.IliveStreams = v
	// fmt.Printf("%+v\n", v)
	return n
}

func (n *nav) CurrentStream(v liveStream) *nav {
	n.IcurrentStream = v
	// fmt.Printf("%+v\n", v)
	return n
}

func (n *nav) OnMount(ctx app.Context) {
	n.isFirstLoad = true
	// url := strings.Split(ctx.Page().URL().Path, "/")
	// n.classSvg = parseSvg(url[1])
}

func (n *nav) OnNav(ctx app.Context) {
	if n.isFirstLoad {
		n.isFirstLoad = false
		ctx.ScrollTo(strings.TrimPrefix(ctx.Page().URL().Path, "/"))
	}
}

func parseTitle(t string) string {
	t = strings.ToLower(t)
	if strings.Contains(t, "wotlk") {
		return wotlkSVG
	}
	if strings.Contains(t, "tbc") {
		return tbcSVG
	}
	return retailSVG
}

func parseSvg(s string) string {
	if s == "death_knight" {
		return dkSVG
	}
	if s == "demon_hunter" {
		return dhSVG
	}
	if s == "druid" {
		return druidSVG
	}
	if s == "hunter" {
		return hunterSVG
	}
	if s == "mage" {
		return mageSVG
	}
	if s == "monk" {
		return monkSVG
	}
	if s == "paladin" {
		return paladinSVG
	}
	if s == "priest" {
		return priestSVG
	}
	if s == "rogue" {
		return rogueSVG
	}
	if s == "shaman" {
		return shamanSVG
	}
	if s == "warlock" {
		return warlockSVG
	}
	if s == "warrior" {
		return warriorSVG
	}
	return retailSVG
}

func (n *nav) Render() app.UI {
	return app.Div().
		Class("nav").
		Class("fill").
		// Class("unselectable").
		Class(n.Iclass).
		Body(
			ui.Stack().
				Class("app-title").
				Class("hspace-out").
				Middle().
				Content(
					app.Header().
						Body(
							app.A().
								Class("view-desktop").
								Class("hApp").
								Class("focus").
								Class("glow").
								Href("/").
								Text("Pvptv.io"),
							app.If(len(n.IcurrentClass) != 0,
								app.If(n.IcurrentClass == "demon_hunter",
									app.A().
										Class("view-mobile").
										Class("hApp").
										Class("focus").
										Class("glow").
										Href("/").
										Text("↩ D. HUNTER"),
								),
								app.If(n.IcurrentClass == "death_knight",
									app.A().
										Class("view-mobile").
										Class("hApp").
										Class("focus").
										Class("glow").
										Href("/").
										Text("↩ D. KNIGHT"),
								),
								app.If(n.IcurrentClass != "death_knight" && n.IcurrentClass != "demon_hunter",
									app.A().
										Class("view-mobile").
										Class("hApp").
										Class("focus").
										Class("glow").
										Href("/").
										Text("↩ "+n.IcurrentClass),
								),
							),
						),
				),
			ui.Stack().
				Center().
				Middle().
				Content(
					app.Div().Body(
						app.Header().
							Class("icon-circle-desktop").
							Class("hApp h2 h2-nav").
							Class("hspace-out").
							// Class("focus").
							Body(
								app.If(n.IcurrentClass == "demon_hunter",
									app.H2().
										Text("Demon Hunter"),
								),
								app.If(n.IcurrentClass == "death_knight",
									app.H2().
										Text("Death Knight"),
								),
								app.If(n.IcurrentClass != "death_knight" && n.IcurrentClass != "demon_hunter",
									app.H2().
										Text(n.IcurrentClass),
								),
							),
					),
				),
			app.Nav().
				Class("nav-content").
				Body(
					app.Div().
						Class("nav-streams-class-selected").
						Body(
							ui.Stack().
								Class("nav-streams-stack").
								Middle().
								Content(
									app.Div().
										Class("hspace-out").
										Body(
											app.Range(n.IliveStreams).Slice(func(i int) app.UI {
												lr := n.IliveStreams[i]
												if lr.Online == true {
													gameVersionIcon := parseTitle(lr.Title)
													return app.Div().Class("stream-label-style").
														Body(
															newLink().
																ID(lr.Slug).
																Class("glow space-label").
																Label(lr.Name ).
																Href("/" + n.IcurrentClass + "/" + lr.Slug).
																Help(lr.Title).
																Icon(newSVGIcon().RawSVG(gameVersionIcon)).
																Focus(lr.Slug == n.IcurrentStream.Slug),
															newLink().
																ID(lr.Slug).
																Class("glow unresponsive").
																Label("🔴 " + strconv.Itoa(lr.Viewers)).
																// Href("/" + n.IcurrentClass + "/" + lr.Slug).
																// Help(lr.Title).
																Focus(lr.Slug == n.IcurrentStream.Slug),

														)
												} else {
													return app.Div().Class("stream-offline").
													Body(
														newLink().
														ID(lr.Slug).
														Class("offlineLink").
														// Icon(newSVGIcon().RawSVG(playSVG)).
														Label(lr.Name).
														Href("/" + n.IcurrentClass + "/" + lr.Slug).
														Focus(lr.Slug == n.IcurrentStream.Slug),
													)
												}
											}),											
										),
								),
						),
					app.Div().
						Class("nav-support").
						Class("hspace-out").
						Body(
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(coffeeSVG)).
								Label("Buy me a coffee").
								Href(buyMeACoffeeURL),
							// newLink().
							// 	Class("glow").
							// 	Icon(newSVGIcon().RawSVG(githubSVG)).
							// 	Label("GitHub").
							// 	Href(githubURL),
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(twitterSVG)).
								Label("Twitter").
								Href(twitterURL),
						),
				),
		)
}
