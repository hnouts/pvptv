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
				Class("support").
				Class("text-center").
				Body(
					ui.Stack().
						Class("sticky-pic-thrall").
						Middle().
						Bottom().
						Content(
							ui.Icon().
								Size(320).
								Src("/web/thrall_dude_paysage.png"),
						),
					ui.Stack().
						Class("sticky-pic-quest").
						Middle().
						Bottom().
						Content(
							ui.Icon().
								Size(50).
								Src("/web/quest-icon.jpg"),
						),
					app.A().
						Class("default").
						Href("https://forms.gle/y5wj532gvtgwWkca6").
						Body(
							app.Header().
								Class("h3").
								Text("Streamers wanted"),
							app.P().
								Class("subtext").
								Text("Share a pvp streamer you like here."),
						),
				),
		)
}
