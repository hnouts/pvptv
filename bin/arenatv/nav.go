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
	showDropdown   bool
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

func parseSpecList(mc string, s []specList, c string) (string, string) {
	// Range through specList, if class = c then spec = current item
	for _, spec := range s {
		if c == spec.SPClass {
			if c == mc {
				return "main", parseSpecToSvg(spec.SPSpec)
			}
			return "alt", parseSpecToSvg(spec.SPSpec)
		}
	}

	return "meta", returnMetaForGivenClass()
}

func returnMetaForGivenClass() string {
	return websiteSVG
}

func parseSpecToSvg(s string) string {
	if svg, ok := specSVGs[s]; ok {
		return svg
	}
	return websiteSVG
}

var specSVGs = map[string]string{
	"mistweaver":        mistweaverMonkSVG,
	"windwalker":        windwalkerMonkSVG,
	"holyPaladin":       holyPaladinSVG,
	"ret":               retPaladinSVG,
	"restorationShaman": restorationShamanSVG,
	"elemental":         elementalShamanSVG,
	"enhancement":       enhancementShamanSVG,
	"restorationDruid":  restorationDruidSVG,
	"balance":           balanceDruidSVG,
	"feral":             feralDruidSVG,
	"shadow":            shadowPriestSVG,
	"disc":              discPriestSVG,
	"holyPriest":        holyPriestSVG,
	"destruction":       destructionWarlockSVG,
	"demonology":        demonologyWarlockSVG,
	"affliction":        afflictionWarlockSVG,
	"arms":              armsWarriorSVG,
	"fury":              furyWarriorSVG,
	"protection":        protectionWarriorSVG,
	"bm":                bmHunterSVG,
	"survival":          survivalHunterSVG,
	"mm":                mmHunterSVG,
	"sub":               subRogueSVG,
	"outlaw":            outlawRogueSVG,
	"assa":              assassinationRogueSVG,
	"fire":              fireMageSVG,
	"frostMage":         frostMageSVG,
	"arcane":            arcaneMageSVG,
	"havoc":             havocDhSVG,
	"vengeance":         vengeanceDhSVG,
	"blood":             bloodDkSVG,
	"unholy":            unholyDkSVG,
	"frostDk":           frostDkSVG,
	"preservation":      preservationEvokerSVG,
	"devastation":       devastationEvokerSVG,
}

func handleClick(ctx app.Context, e app.Event) {
	ctx.Reload()
}
func (n *nav) handleMouseOver(ctx app.Context, e app.Event) {
	n.showDropdown = true
}

func (n *nav) handleMouseOut(ctx app.Context, e app.Event) {
	n.showDropdown = false
}

func (n *nav) renderDropdown() app.UI {
	if !n.showDropdown {
		return nil
	}

	links := make([]app.UI, 0)
	for _, c := range getAllClasses() {
		link := ui.Link().
			Class("link heading fit icon-circle").
			Href("/" + c.Slug).
			Label(c.Name).
			Icon(c.Svg).
			OnClick(handleClick)

		links = append(links, app.Div().Body(link))
	}

	return app.Div().
		Class("dropdown-container").
		Body(
			app.Div().
				Class("dropdown").
				Body(
					app.Div().
						Body(links...),
				),
		)
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
								Text("↩ HOME"),
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
					app.Div().
						OnMouseOver(n.handleMouseOver).
						Body(
							app.Header().
								Class("icon-circle-desktop").
								Class("hApp h2 h2-nav").
								Class("hspace-out").
								Style("cursor", "pointer").
								Body(
									app.H2().
										Text(n.IcurrentClass).
										ID("header-id"),
									n.renderDropdown(),
								),
						),
				),

			app.Nav().
				Class("nav-content").
				Body(
					app.Div().
						OnMouseOut(n.handleMouseOut).
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
												Text("Streaming"),
											app.Range(n.IliveStreams).Slice(func(i int) app.UI {
												lr := n.IliveStreams[i]
												if lr.Online && lr.GameName == "World of Warcraft" {
													opt, specIcon := parseSpecList(lr.MainClass, lr.SpecList, n.IcurrentClass)
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
												if !lr.Online || lr.GameName != "World of Warcraft" {
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
							newLink().
								Class("glow").
								Icon(newSVGIcon().RawSVG(githubSVG)).
								Label("GitHub").
								Href(githubURL),
						),
				),
		)
}
