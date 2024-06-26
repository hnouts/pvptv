package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

const (
	loaderSize          = 78
	controlIconSize     = 18
	controlMainIconSize = 30
)

type twitchState int

const (
	unstarted = -1
	ended     = 0
	playing   = 1
	paused    = 2
	buffering = 3
	videoCued = 5
)

type twitchPlayer struct {
	app.Compo

	Iclass            string
	Istream           liveStream
	IonPlaybackChange func(app.Context, bool)

	initPlayer  sync.Once
	stream      liveStream
	player      app.Value
	isPlaying   bool
	isBuffering bool
	canBack     bool
	volume      volume
	err         error

	realeaseOnReady func()
	releaseOnError  func()
}

func newTwitchPlayer() *twitchPlayer {
	return &twitchPlayer{}
}

func (p *twitchPlayer) Class(v string) *twitchPlayer {
	if v == "" {
		return p
	}
	if p.Iclass != "" {
		p.Iclass += " "
	}
	p.Iclass += v
	return p
}

func (p *twitchPlayer) Stream(v liveStream) *twitchPlayer {
	p.Istream = v
	return p
}

func (p *twitchPlayer) OnPlaybackChange(v func(app.Context, bool)) *twitchPlayer {
	p.IonPlaybackChange = v
	return p
}

func (p *twitchPlayer) OnMount(ctx app.Context) {
	p.isBuffering = true
	p.volume = loadVolume(ctx)
}

func (p *twitchPlayer) OnNav(ctx app.Context) {
	p.canBack = app.Window().Get("history").Get("length").Int() > 1
}

func (p *twitchPlayer) OnDismount() {
	if p.realeaseOnReady != nil {
		p.realeaseOnReady()
	}
	if p.releaseOnError != nil {
		p.releaseOnError()
	}
	p.player = nil
}

func (p *twitchPlayer) OnResize(ctx app.Context) {
}

func (p *twitchPlayer) loadVideo(ctx app.Context) {

	if p.Istream.Slug != p.stream.Slug {
		p.stream = p.Istream
		if p.player != nil {
			p.loadChannel(ctx, p.stream.twitchSlug())
			return
		}
	}

	p.initPlayer.Do(func() {
		READY := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
			ctx.Dispatch(func(ctx app.Context) {
				p.setVolume(ctx, p.volume.Value)
				p.play(ctx)
			})
			return nil
		})
		p.realeaseOnReady = READY.Release

		ERROR := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
			ctx.Dispatch(func(ctx app.Context) {
				p.onError(ctx, args)
			})
			return nil
		})
		p.releaseOnError = ERROR.Release

		p.player = app.Window().
			Get("Twitch").
			Get("Player").
			New("twitch-player", map[string]interface{}{
				"channel":         p.stream.twitchSlug(),
				"parent":          "localhost",
				"allowfullscreen": true,
			})

		p.player.Call("addEventListener", "ready", READY)
		p.player.Call("addEventListener", "error", ERROR)
	})
}

func (p *twitchPlayer) onError(ctx app.Context, args []app.Value) {
	fmt.Println("data Err:", ctx)
	code := args[0].Get("data").Int()
	msg := ""

	switch code {
	case 2:
		msg = "invalid video parameter values"

	case 5:
		msg = "loading video failed"

	case 100:
		msg = "video not found"

	case 101, 150:
		msg = "video cannot be played in embedded players"

	default:
		msg = "unkown error"

	}

	p.err = errors.New("youtube player error").
		Tag("code", code).
		Tag("description", msg)

	fmt.Println("error:", p.err)
}

func (p *twitchPlayer) OnUpdate(ctx app.Context) {
	if p.Istream.Slug != "" && p.stream.Slug != p.Istream.Slug {
		p.loadVideo(ctx)
	}
}

func (p *twitchPlayer) Render() app.UI {
	return app.Div().
		Class("youtube").
		Class(p.Iclass).
		Body(
			app.Div().
				Class("youtube-video").
				Body(
					app.Div().
						ID("twitch-player").
						Body(
							app.Script().Src("https://player.twitch.tv/js/embed/v1.js")),
				),
			app.If(p.isBuffering || p.err != nil,
				app.Div().
					Class("youtube-noplay").
					Class("fill").
					Class("background-overlay").
					Body(
						newLoader().
							Class("hspace-out").
							Class("vspace-stretch").
							Size(loaderSize).
							Title(p.stream.Name).
							PlayerClasses(p.stream.ClassList).
							Loading(p.isBuffering).
							Err(p.err),
					),
			),
		)
}

func (p *twitchPlayer) onPlayClicked(ctx app.Context, e app.Event) {
	p.play(ctx)
}

func (p *twitchPlayer) onPauseClicked(ctx app.Context, e app.Event) {
	p.pause(ctx)
}

func (p *twitchPlayer) onBackClicked(ctx app.Context, e app.Event) {
	app.Window().Get("history").Call("back")
}

func (p *twitchPlayer) onSoundClicked(ctx app.Context, e app.Event) {
	if p.volume.Value == 0 {
		p.setVolume(ctx, p.volume.LastHearableValue)
		return
	}
	p.setVolume(ctx, 0)
}

func (p *twitchPlayer) onVolumeChanged(ctx app.Context, e app.Event) {
	volume, _ := strconv.Atoi(ctx.JSSrc().Get("value").String())
	p.setVolume(ctx, volume)
}

func (p *twitchPlayer) loadChannel(ctx app.Context, id string) {
	p.player.Call("setChannel", id)
}

func (p *twitchPlayer) play(ctx app.Context) {
	println("Starting streaming...")
	p.player.Call("play")
	p.isPlaying = true
	p.isBuffering = false

}

func (p *twitchPlayer) pause(ctx app.Context) {
	p.player.Call("pause")
	p.isPlaying = false
}

func (p *twitchPlayer) setVolume(ctx app.Context, v int) {
	var floatVol float64 = float64(v)
	if v == 0 {
		p.player.Call("setMuted", true)
	} else {
		p.volume.LastHearableValue = v
		p.player.Call("setMuted", false)
		floatVol = floatVol / 100.0
		p.player.Call("setVolume", math.Floor(floatVol*100)/100)
	}

	p.volume.Value = v
	saveVolume(ctx, p.volume)
}

type volume struct {
	Value             int
	LastHearableValue int
}

func saveVolume(ctx app.Context, v volume) {
	if err := ctx.LocalStorage().Set("player-volume", v); err != nil {
		app.Log(errors.New("saving volume values in local storage failed").
			Wrap(err))
	}
}

func loadVolume(ctx app.Context) volume {
	var v volume
	if err := ctx.LocalStorage().Get("player-volume", &v); err != nil {
		app.Log(errors.New("saving player status in local storage failed").
			Wrap(err))
	}

	if v == (volume{}) {
		v.Value = 50
		v.LastHearableValue = 50
	}
	return v
}
