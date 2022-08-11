package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type githubSponsor struct {
	app.Compo

	Iclass string
}

func newGithubSponsor() *githubSponsor {
	return &githubSponsor{}
}

func (s *githubSponsor) Class(v string) *githubSponsor {
	s.Iclass = app.AppendClass(s.Iclass, v)
	return s
}

func (s *githubSponsor) Render() app.UI {
	return ui.Stack().
		Class(s.Iclass).
		Center().
		Middle().
		Content(
			app.Aside().
				Class("magnify").
				Class("text-center").
				// Body(
				// 	ui.Flow().
				// 		// StretchItems().
				// 		// Spacing(84).
				// 		Content(
				// 			ui.Stack().
				// 				Bottom().
				// 				Content(
				// 					ui.Icon().
				// 						// Class("unselectable").
				// 						Size(320).
				// 						Src("/web/thrall_dude_paysage.png"),
				// 				),
				// 		),
				// ),
				Body(
					// ui.Stack().
					// 	Content(
					ui.Stack().
						Class("sticky-pic").
						Middle().
						Bottom().
						Content(
							ui.Icon().
								// Class("cover_thrall").
								Size(320).
								Src("/web/thrall_dude_paysage.png"),
						),
					// ),
					app.A().
						Class("default").
						Href("localhost").
						Body(
							app.Header().
								Class("h3").
								Class("default").
								Text("Help me find streamers"),
							app.P().
								Class("subtext").
								Text("Add a pvp streamer you like here."),
						),
				),
		)
}
