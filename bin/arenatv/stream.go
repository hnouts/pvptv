package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
	"golang.org/x/exp/rand"
)

const (
	menuWidth = 282
)

type stream struct {
	app.Compo

	lives             []liveStream
	current           liveStream
	isPlaying         bool
	isUpdateAvailable bool
	randHistory       []liveStream
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
	rand.Seed(uint64(time.Now().UnixNano()))
	url := strings.Split(ctx.Page().URL().Path, "/")
	allLives := getLiveStreamers()

	for _, streamer := range allLives {
		if isCurrentClass(streamer, url[1]) {
			r.lives = append(r.lives, streamer)
		}
	}

	fmt.Printf("%+v\n", r.lives)
}

func isCurrentClass(streamer liveStream, class string) bool {

	return streamer.Class == class
}

func (r *stream) load(ctx app.Context) {
	url := strings.Split(ctx.Page().URL().Path, "/")
	slug := strings.TrimPrefix(ctx.Page().URL().Path, "/"+url[1]+"/")

	if slug == "" {
		println("NO SLUG? PICK RANDOM")
		r.current = r.randomStreamer()
		u := *ctx.Page().URL()
		u.Path = "/" + r.current.Class + "/" + r.current.Slug
		ctx.Page().ReplaceURL(&u)
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
	// ctx.Page().SetDescription(fmt.Sprintf("Listen to Lo-fi music stream %s on the Lofimusic open-source player: an installable Progressive Web app (PWA) written in Go (Golang).", r.current.Name))
	// ctx.Page().SetImage("https://lofimusic.app/web/covers/" + slug + ".png")
}

func (r *stream) OnAppUpdate(ctx app.Context) {
	r.isUpdateAvailable = ctx.AppUpdateAvailable()
}

func (r *stream) randomStreamer() liveStream {
	if len(r.randHistory) == 0 {
		r.randHistory = make([]liveStream, len(r.lives))
		copy(r.randHistory, r.lives)
	}

	idx := rand.Intn(len(r.randHistory))
	stream := r.randHistory[idx]

	copy(r.randHistory[idx:], r.randHistory[idx+1:])
	r.randHistory = r.randHistory[:len(r.randHistory)-1]

	return stream
}

func (r *stream) Render() app.UI {
	return app.Main().
		Class("stream").
		Class("fill").
		Body(
			// newYouTubePlayer().
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
					CurrentClass(r.current.Class).
					LiveStreams(r.lives).
					CurrentStream(r.current)).
				HamburgerMenu(newNav().
					Class("background-overlay").
					LiveStreams(r.lives).
					CurrentStream(r.current)).
				Content(
					app.Aside().
						Class("stream-update").
						Class("app-title").
						Class("hspace-out").
						Body(
							ui.Stack().
								Class("fill").
								Right().
								Middle().
								Content(
									app.If(r.isUpdateAvailable,
										newLink().
											Class("link-update").
											Class("glow").
											Label("Update").
											Icon(newSVGIcon().RawSVG(downloadSVG)).
											OnClick(r.onUpdateClick),
									),
								),
						),
					app.Div().
						Class("hspace-out").
						Class("vspace-content").
						Body(
							newInfo().
								Stream(r.current).
								Playing(r.isPlaying),
						),
				),
		)
}

func (r *stream) onPlaybackChange(ctx app.Context, isPlaying bool) {
	r.isPlaying = isPlaying
}

func (r *stream) onUpdateClick(ctx app.Context) {
	ctx.Reload()
}
