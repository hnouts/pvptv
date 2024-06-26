package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

const (
	menuWidth = 282
)

type stream struct {
	app.Compo

	lives             []liveStream
	current           liveStream
	classInUrl        string
	isPlaying         bool
	isUpdateAvailable bool
}

func newStream() *stream {
	return &stream{}
}

func (r *stream) OnPreRender(ctx app.Context) {
	r.init(ctx)
	r.load(ctx)
}

func (r *stream) OnMount(ctx app.Context) {
	r.init(ctx)
}

func (r *stream) OnNav(ctx app.Context) {
	r.load(ctx)
}

func (r *stream) OnResize(ctx app.Context) {
	r.ResizeContent()
}

func (r *stream) init(ctx app.Context) {
	defer timer()()
	var twitchToken twitchAuth = <-getTwitchToken()
	url := strings.Split(ctx.Page().URL().Path, "/")
	currentPageLives := getLiveStreamersForAClass(url[1])
	var livesData streamData = <-getMultipleStreamsData(currentPageLives, twitchToken.AccessToken)

	if len(livesData.Data) == 0 {
		for _, s := range currentPageLives {
			s.Online = false
			r.lives = append(r.lives, s)
		}
	} else {
		for _, cpl := range currentPageLives {
			for _, ld := range livesData.Data {
				if cpl.Slug == ld.UserLogin {
					cpl.Online = true
					cpl.Title = ld.Title
					cpl.Viewers = ld.ViewerCount
					cpl.GameName = ld.GameName
				}
			}
			r.lives = append(r.lives, cpl)
		}
	}

	sort.SliceStable(r.lives, func(i, j int) bool {
		return r.lives[i].Viewers > r.lives[j].Viewers
	})

	sort.SliceStable(r.lives, func(i, j int) bool {
		return strings.Compare(r.lives[i].GameName, r.lives[j].GameName) > 0
	})
}

func (r *stream) load(ctx app.Context) {
	url := strings.Split(ctx.Page().URL().Path, "/")
	r.classInUrl = url[1]
	slug := strings.TrimPrefix(ctx.Page().URL().Path, "/"+url[1]+"/")
	if slug == "" || len(url) < 3 {
		r.current = r.lives[0]
	} else {
		for _, lr := range r.lives {
			if slug == lr.Slug {
				r.current = lr
				break
			}
		}
	}

	r.isUpdateAvailable = ctx.AppUpdateAvailable()
	r.isPlaying = false

	ctx.Page().SetTitle(fmt.Sprintf("%s Stream", r.current.Name))
	ctx.Page().SetDescription(fmt.Sprintf("Watch %s on PvPtv.io: an installable Progressive Web app (PWA) written in Go (Golang).", r.current.Name))
}

func (r *stream) OnAppUpdate(ctx app.Context) {
	r.isUpdateAvailable = ctx.AppUpdateAvailable()
}

func (r *stream) Render() app.UI {
	return app.Main().
		Class("stream").
		Class("fill").
		Body(
			newTwitchPlayer().
				Class("stream-player").
				Class("fill").
				Stream(r.current).
				OnPlaybackChange(r.onPlaybackChange),
			ui.Shell().
				Class("stream-shell").
				Class("fill").
				PaneWidth(menuWidth).
				Menu(newNav().
					CurrentClass(r.classInUrl).
					LiveStreams(r.lives).
					CurrentStream(r.current)).
				HamburgerMenu(newNav().
					Class("background-overlay").
					CurrentClass(r.classInUrl).
					LiveStreams(r.lives).
					CurrentStream(r.current)).
				Content(
					app.If(!r.current.Online,
						app.Div().
							Class("hspace-out").
							Class("vspace-content").
							Body(
								newInfo().
									Stream(r.current).
									Playing(r.isPlaying),
							),
					),
				),
		)
}

func (r *stream) onPlaybackChange(ctx app.Context, isPlaying bool) {
	r.isPlaying = isPlaying
}
