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
	isFirstLoad    bool
}

func handleMouseOver(ctx app.Context, e app.Event) {
	dropdown := app.Window().GetElementByID("dropdown-id")
	if dropdown != nil {
		dropdown.Set("style", "display:block;")
	}
}

func handleMouseOut(ctx app.Context, e app.Event) {
	dropdown := app.Window().GetElementByID("dropdown-id")
	if dropdown != nil {
		dropdown.Set("style", "display:none;")
		header := app.Window().GetElementByID("header-id")
		if header != nil {
			header.Set("innerText","test")
		}
	}
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
		return "wotlk"
	}
	if strings.Contains(t, "classic") {
		return "wotlk"
	}
	if strings.Contains(t, "80") {
		return "wotlk"
	}
	if strings.Contains(t, "wrath") {
		return "wotlk"
	}
	if strings.Contains(t, "lich") {
		return "wotlk"
	}

	return "retail"
}

func parseSpecList(version string, mc string, s []specList, c string) (string, string) {
	// range speclist, if class = c then spec = current item

	if c == mc {
		for _, spec := range s {
			if c == spec.SPClass {
				return "main", parseSpecToSvg(spec.SPSpec)
			}
		}
	} else {
		for _, spec := range s {
			if c == spec.SPClass {
				return "alt", parseSpecToSvg(spec.SPSpec)
			}
		}
	}

	return "meta", returnMetaforGivenClass(version, c)
}

// func isClassicStreaming(stream []liveStream) bool {
// 	for _, s := range stream {
// 		if s.Online == true && parseTitle(s.Title) == "wotlk" && s.GameName == "World of Warcraft" {
// 			return true
// 		}
// 	}
// 	return false
// }

// func isRetailStreaming(stream []liveStream) bool {
// 	for _, s := range stream {
// 		if s.Online == true && parseTitle(s.Title) == "retail" && s.GameName == "World of Warcraft" {
// 			return true
// 		}
// 	}
// 	return false
// }

func returnMetaforGivenClass(version string, class string) string {
	// LAST UPDATE SHADOWLAND SEASON 4 - SEPTEMBER 2022

	if version == "classic" {
		// switch class {
		// case "death_knight":
		// 	return unholyDkSVG
		// case "hunter":
		// 	return mmHunterSVG
		// case "mage":
		// 	return frostMageSVG
		// case "rogue":
		// 	return subRogueSVG
		// case "shaman":
		// 	return elementalShamanSVG
		// case "warlock":
		// 	return destructionWarlockSVG
		// case "warrior":
		// 	return armsWarriorSVG
		// }
		return websiteSVG
	} else {
		// switch class {
		// case "demon_hunter":
		// 	return havocDhSVG
		// case "death_knight":
		// 	return unholyDkSVG
		// case "hunter":
		// 	return mmHunterSVG
		// case "mage":
		// 	return fireMageSVG
		// case "monk":
		// 	return windwalkerMonkSVG
		// case "priest":
		// 	return discPriestSVG
		// case "rogue":
		// 	return subRogueSVG
		// case "shaman":
		// 	return elementalShamanSVG
		// case "warlock":
		// 	return destructionWarlockSVG
		// case "warrior":
		// 	return armsWarriorSVG
		// }
		return websiteSVG
	}
}

func parseSpecToSvg(s string) string {
	switch s {
	case "mistweaver":
		return mistweaverMonkSVG
	case "windwalker":
		return windwalkerMonkSVG
	case "holyPaladin":
		return holyPaladinSVG
	case "ret":
		return retPaladinSVG
	case "restorationShaman":
		return restorationShamanSVG
	case "elemental":
		return elementalShamanSVG
	case "enhancement":
		return enhancementShamanSVG
	case "restorationDruid":
		return restorationDruidSVG
	case "balance":
		return balanceDruidSVG
	case "feral":
		return feralDruidSVG
	case "shadow":
		return shadowPriestSVG
	case "disc":
		return discPriestSVG
	case "holyPriest":
		return holyPriestSVG
	case "destruction":
		return destructionWarlockSVG
	case "demonology":
		return demonologyWarlockSVG
	case "affliction":
		return afflictionWarlockSVG
	case "arms":
		return armsWarriorSVG
	case "fury":
		return furyWarriorSVG
	case "protection":
		return protectionWarriorSVG
	case "bm":
		return bmHunterSVG
	case "survival":
		return survivalHunterSVG
	case "mm":
		return mmHunterSVG
	case "sub":
		return subRogueSVG
	case "outlaw":
		return outlawRogueSVG
	case "assa":
		return assassinationRogueSVG
	case "fire":
		return fireMageSVG
	case "frostMage":
		return frostMageSVG
	case "arcane":
		return arcaneMageSVG
	case "havoc":
		return havocDhSVG
	case "vengeance":
		return vengeanceDhSVG
	case "blood":
		return bloodDkSVG
	case "unholy":
		return unholyDkSVG
	case "frostDk":
		return frostDkSVG
	case "preservation":
		return preservationEvokerSVG
	case "devastation":
		return devastationEvokerSVG
	default:
		return websiteSVG
	}
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
								Text("PvPtv.io"),
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
					// app.Div().Body(
					// 	app.Header().
					// 	Class("icon-circle-desktop").
					// 	Class("hApp h2 h2-nav").
					// 	Class("hspace-out").
					// 	Body(
					// 		app.If(n.IcurrentClass == "demon_hunter",
					// 			app.H2().
					// 				Text("Demon Hunter"),
					// 		),
					// 		app.If(n.IcurrentClass == "death_knight",
					// 			app.H2().
					// 				Text("Death Knight"),
					// 		),
					// 		app.If(n.IcurrentClass != "death_knight" && n.IcurrentClass != "demon_hunter",
					// 			app.H2().
					// 				Text(n.IcurrentClass).
					// 				ID("header-id").
					// 				OnMouseOver(handleMouseOver).
					// 				OnMouseOut(handleMouseOut).
					// 				Body(
					// 					app.Div().
					// 						Class("dropdown").
					// 						ID("dropdown-id").
					// 						Style("display", "none").
					// 						Body(
					// 							app.Ul().
					// 								Class("dropdown-menu").
					// 								Attr("aria-labelledby", "dropdownMenuButton").
					// 								Body(
					// 									app.Li().Text("Class 1"),
					// 									app.Li().Text("Class 2"),
					// 									app.Li().Text("Class 3"),
					// 								),
					// 						),
					// 				),
					// 		),
					// 	),
					// ),
					
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
											app.Div().Class("separator"),
											app.H3().
												Class("hApp h3").
												Text("Classic"),
											app.Range(n.IliveStreams).Slice(func(i int) app.UI {
												lr := n.IliveStreams[i]
												if lr.Online == true && parseTitle(lr.Title) == "wotlk" && lr.GameName == "World of Warcraft" {

													opt, specIcon := parseSpecList("classic", lr.MainClass, lr.SpecList, n.IcurrentClass)

													return app.Div().Class("stream-label-style").
														Body(
															newLink().
																ID(lr.Slug).
																Class("glow icon-circle "+opt).
																Label(lr.Name).
																Href("/"+n.IcurrentClass+"/"+lr.Slug).
																Help(lr.Title).
																Icon(newSVGIcon().RawSVG(specIcon)).
																Focus(lr.Slug == n.IcurrentStream.Slug),
															newLink().
																ID(lr.Slug).
																Class("glow unresponsive spaced-on-mobile").
																Label("🔴 "+strconv.Itoa(lr.Viewers)).
																Focus(lr.Slug == n.IcurrentStream.Slug),
														)
												} else {
													return app.Div()
												}
											}),
											app.Div().Class("separator"),
											app.H3().
												Class("hApp h3").
												Text("Retail"),
											app.Range(n.IliveStreams).Slice(func(i int) app.UI {
												lr := n.IliveStreams[i]
												if lr.Online == true && parseTitle(lr.Title) == "retail" && lr.GameName == "World of Warcraft" {
													opt, specIcon := parseSpecList("retail", lr.MainClass, lr.SpecList, n.IcurrentClass)
													return app.Div().Class("stream-label-style").
														Body(
															newLink().
																ID(lr.Slug).
																Class("glow icon-circle "+opt).
																Label(lr.Name).
																Href("/"+n.IcurrentClass+"/"+lr.Slug).
																Help(lr.Title).
																Icon(newSVGIcon().RawSVG(specIcon)).
																Focus(lr.Slug == n.IcurrentStream.Slug),
															newLink().
																ID(lr.Slug).
																Class("glow unresponsive spaced-on-mobile").
																Label("🔴 "+strconv.Itoa(lr.Viewers)).
																Focus(lr.Slug == n.IcurrentStream.Slug),
														)
												} else {
													return app.Div()
												}
											}),

											app.Div().Class("separator"),
											app.H3().
												Class("hApp h3").
												Text("Offline"),
											app.Range(n.IliveStreams).Slice(func(i int) app.UI {
												lr := n.IliveStreams[i]
												if lr.Online == false || lr.GameName != "World of Warcraft" {
													return app.Div().Class("stream-offline").
														Body(
															newLink().
																ID(lr.Slug).
																Class("offlineLink").
																Label(lr.Name).
																Href("/" + n.IcurrentClass + "/" + lr.Slug).
																Focus(lr.Slug == n.IcurrentStream.Slug),
														)
												} else {
													return app.Div()
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
