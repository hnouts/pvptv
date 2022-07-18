package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
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

	realeaseOnReady      func()
	releaseOnStateChange func()
	releaseOnError       func()
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
	if p.releaseOnStateChange != nil {
		p.releaseOnStateChange()
	}
	if p.releaseOnError != nil {
		p.releaseOnError()
	}
	p.player = nil
}

func (p *twitchPlayer) OnResize(ctx app.Context) {
}

func (p *twitchPlayer) loadVideo(ctx app.Context) {

	// if isOnYouTubeIframeAPIReady := app.Window().Get("isOnYouTubeIframeAPIReady").Bool(); !isOnYouTubeIframeAPIReady && app.IsClient {
	// 	fmt.Println("CALLED isOnYouTubeIframeAPIReady :", ctx)
	// 	ctx.Async(func() {
	// 		time.Sleep(time.Millisecond * 1000)
	// 		ctx.Dispatch(p.loadVideo)
	// 	})
	// 	return
	// }

	if p.Istream.Slug != p.stream.Slug {
		p.stream = p.Istream
		if p.player != nil {
			p.loadVideoByID(ctx, p.stream.twitchID())
			return
		}
	}

	p.initPlayer.Do(func() {
		// onReady := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		// 	ctx.Dispatch(func(ctx app.Context) {
		// 		p.setVolume(ctx, p.volume.Value)
		// 		p.play(ctx)
		// 	})
		// 	return nil
		// })
		// p.realeaseOnReady = onReady.Release

		// onStateChange := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		// 	ctx.Dispatch(func(ctx app.Context) {
		// 		p.onStateChange(ctx, args)
		// 	})
		// 	return nil
		// })
		// p.releaseOnStateChange = onStateChange.Release

		// onError := app.FuncOf(func(this app.Value, args []app.Value) interface{} {
		// 	ctx.Dispatch(func(ctx app.Context) {
		// 		p.onError(ctx, args)
		// 	})
		// 	return nil
		// })
		// p.releaseOnError = onError.Release

		// p.player = app.Window().
		// 	Get("YT").
		// 	Get("Player").
		// 	New("youtube-player", map[string]interface{}{
		// 		"channel": p.stream.twitchID(),
		// 		"parent":  "localhost",
		// 		"events": map[string]interface{}{
		// 			"onReady":       onReady,
		// 			"onStateChange": onStateChange,
		// 			"onError":       onError,
		// 		},
		// 	})
		// fmt.Println("P.PLAYER:")
		// fmt.Printf("%+v\n", p)

	})
}

func (p *twitchPlayer) onStateChange(ctx app.Context, args []app.Value) {
	switch args[0].Get("data").Int() {
	case unstarted:
		p.isPlaying = false
		p.isBuffering = false

	case ended:
		p.isPlaying = false
		p.isBuffering = false
		if p.err == nil {
			p.play(ctx)
		}

	case playing:
		p.isPlaying = true
		p.isBuffering = false
		p.err = nil

	case paused:
		p.isPlaying = false
		p.isBuffering = false

	case buffering:
		p.isBuffering = true
		p.err = nil
	}

	if p.IonPlaybackChange != nil {
		ctx.Emit(func() {
			p.IonPlaybackChange(ctx, p.isPlaying)
		})
	}
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
	fmt.Println("+ :", p.Istream.Slug)
	if p.Istream.Slug != "" && p.stream.Slug != p.Istream.Slug {
		p.loadVideo(ctx)
	}
}

func (p *twitchPlayer) Render() app.UI {
	volumeDisplay := ""
	if p.player == nil {
		volumeDisplay = "disabled"
	}

	fmt.Println("stream:", p.stream.Slug)

	return app.Div().
		Class("youtube").
		Class(p.Iclass).
		Body(
			app.Div().
				Class("youtube-video").
				Body(
					app.Div().
						ID("youtube-player").
						Class("unselectable").
						Body(
							// app.Script().
							// app.Script().Src("https://player.twitch.tv/js/embed/v1.js"),
							// app.Script().Src("https://www.youtube.com/iframe_api"),
							// Async(true),
							app.If(p.stream.Slug != "",
								app.IFrame().
									ID("yt-container").
									Allow("autoplay").
									Allow("accelerometer").
									Allow("encrypted-media").
									Allow("picture-in-picture").
									Sandbox("allow-presentation allow-same-origin allow-scripts allow-popups").
									Src("https://player.twitch.tv/?channel="+p.stream.Slug+"&parent=localhost"),
							)),
				),
			app.If(p.stream.Slug == "" || p.err != nil,
				app.Div().
					Class("youtube-noplay").
					Class("fill").
					Class("background-overlay").
					Body(
						newLoader().
							Class("hspace-out").
							Class("vspace-stretch").
							Size(loaderSize).
							Title("Buffering").
							Description(p.stream.Name).
							Loading(p.isBuffering).
							Err(p.err),
					),
			),
			ui.Stack().
				Class("youtube-controls").
				Class("hspace-out").
				Class("vspace-top").
				Class("vspace-bottom").
				Center().
				Middle().
				Content(
					app.Div().Class("youtube-left-space"),
					newControl().
						Class("youtube-back").
						Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(backwardSVG)).
						Disabled(!p.canBack).
						OnClick(p.onBackClicked),
					app.If(p.isPlaying || p.isBuffering,
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(pauseSVG)).
							Disabled(p.player == nil).
							OnClick(p.onPauseClicked),
					).Else(
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(playSVG)).
							Disabled(p.player == nil || p.isBuffering).
							OnClick(p.onPlayClicked),
					),
					newControl().
						Class("control-main").
						Icon(newSVGIcon().
							Size(controlMainIconSize).
							RawSVG(shuffleSVG)).
						OnClick(p.onShuffleClicked),
					app.If(p.volume.Value > 60,
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(soundHighSVG)).
							Disabled(p.player == nil).
							OnClick(p.onSoundClicked),
					).ElseIf(p.volume.Value > 20,
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(soundMediumSVG)).
							Disabled(p.player == nil).
							OnClick(p.onSoundClicked),
					).ElseIf(p.volume.Value > 0,
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(soundLowSVG)).
							Disabled(p.player == nil).
							OnClick(p.onSoundClicked),
					).Else(
						newControl().Icon(newSVGIcon().
							Size(controlIconSize).
							RawSVG(soundMutedSVG)).
							Disabled(p.player == nil).
							OnClick(p.onSoundClicked),
					),
					app.Div().
						Class("youtube-volume").
						Class(volumeDisplay).
						Body(
							app.Input().
								ID("youtube-volume").
								Type("range").
								Min("0").
								Max("100").
								Value(p.volume.Value).
								OnChange(p.onVolumeChanged).
								OnInput(p.onVolumeChanged),
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

func (p *twitchPlayer) onShuffleClicked(ctx app.Context, e app.Event) {
	ctx.Navigate("/")
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

func (p *twitchPlayer) loadVideoByID(ctx app.Context, id string) {
	p.player.Call("loadVideoById", id, 0)
}

func (p *twitchPlayer) play(ctx app.Context) {
	p.player.Call("playVideo")
}

func (p *twitchPlayer) pause(ctx app.Context) {
	p.player.Call("pauseVideo")
}

func (p *twitchPlayer) setVolume(ctx app.Context, v int) {
	if v == 0 {
		p.player.Call("mute")
	} else {
		p.volume.LastHearableValue = v
		p.player.Call("unMute")
		p.player.Call("setVolume", v)
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
		v.Value = 100
		v.LastHearableValue = 100
	}
	return v
}
