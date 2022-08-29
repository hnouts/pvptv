package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"log"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/cli"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
	"github.com/maxence-charriere/go-app/v9/pkg/logs"
)

const (
	backgroundColor = "#000000"

	buyMeACoffeeURL = "https://www.buymeacoffee.com/hugodev"
	githubURL       = "https://github.com/selkal"
	twitterURL      = "https://twitter.com/hnouts_dev"
)

type options struct {
	Port int `env:"PORT" help:"The port used to listen connections."`
}

type githubOptions struct {
	Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
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
		LoadingLabel: "WoW Arena stream gallery",
		Name:         "Pvptv",
		Image:        "/web/logo.png",
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
	// cli.Register("local").
	// 	Help(`Launches a server that serves the documentation app in a local environment.`).
	// 	Options(&opts)

	// githubOpts := githubOptions{}
	// cli.Register("github").
	// 	Help(`Generates the required resources to run Lofimusic app on GitHub Pages.`).
	// 	Options(&githubOpts)
	// cli.Load()

	// switch cli.Load() {
	// case "local":
	// 	runLocal(ctx, &h, opts)

	// case "github":
	// 	generateGitHubPages(ctx, &h, githubOpts)
	// }
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

func generateGitHubPages(ctx context.Context, h *app.Handler, opts githubOptions) {
	radios := getLiveStreamers()
	slugs := make([]string, len(radios))
	for i, r := range radios {
		slugs[i] = r.Slug
	}

	if err := app.GenerateStaticWebsite(opts.Output, h, slugs...); err != nil {
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
