package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/cli"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
	"github.com/maxence-charriere/go-app/v9/pkg/logs"
)

const (
	backgroundColor = "#000000"

	buyMeACoffeeURL = "https://www.buymeacoffee.com/hugodev"
	githubURL       = "https://github.com/selkal"
	twitterURL      = "https://twitter.com/selkal_dev"
)

type options struct {
	Port int `env:"PORT" help:"The port used to listen connections."`
}

type githubOptions struct {
	Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
	analytics.Add(analytics.NewGoogleAnalytics())
	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	for _, l := range getLiveStreamers() {
		// app.Route("/"+l.Slug, newStream())
		if len(l.ClassList) > 1 {
			for c := range l.ClassList {
				app.Route("/"+l.ClassList[c]+"/"+l.Slug, newStream())
			}
		} else {
			app.Route("/"+l.ClassList[0]+"/"+l.Slug, newStream())
		}
	}

	for _, c := range getAllClasses() {
		app.Route("/"+c.Slug, newStream())
	}
	app.Route("/", newHomePage())
	app.Handle(installApp, handleAppInstall)
	app.Handle(updateApp, handleAppUpdate)
	app.RunWhenOnBrowser()

	defer cancel()
	defer exit()

	h := app.Handler{
		Author:          "Hugo Nouts",
		BackgroundColor: backgroundColor,
		Description:     "WoW arena stream gallery",
		Icon: app.Icon{
			Default: "/web/logo.png",
		},
		Keywords: []string{
			"wow",
			"pvp",
			"arena",
			"stream",
			"2s",
			"3s",
			"battleground",
			"arenas",
			"world of warcraft",
			"warcraft",
			"tbc",
			"wotlk",
			"classic",
			"shadowlands",
			"dragonflight",
		},
		LoadingLabel: "WoW arena stream gallery",
		Name:         "Pvptv",
		Image:        "/web/logo.png",
		RawHeaders: []string{
			// `<script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-1013306768105236" crossorigin="anonymous"></script>`,
			analytics.GoogleAnalyticsHeader("G-1R6LSHKGJZ"),
		},
		Styles: []string{
			"https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
			"/web/arenatv.css",
		},
		ThemeColor: backgroundColor,
		Title:      "Pvptv",
	}
	opts := options{Port: 8080}
	runLocal(ctx, &h, opts)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func runLocal(ctx context.Context, h http.Handler, opts options) {
	app.Logf("%s", logs.New("starting pvptv app server").
		Tag("port", opts.Port),
	)

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", opts.Port),
		Handler: h,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func exit() {
	err := recover()
	if err != nil {
		app.Logf("command failed: %s", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
