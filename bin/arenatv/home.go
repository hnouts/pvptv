package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/analytics"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type homePage struct {
	app.Compo
}

func newHomePage() *homePage {
	return &homePage{}
}

func (p *homePage) OnPreRender(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *homePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("test")
	ctx.Page().SetDescription("test")
	analytics.Page("home", nil)
}

func (p *homePage) Render() app.UI {
	return newPage().
		Index(
			newIndexLink().Title("What is go-app?"),
		).
		Content(
			ui.Flow().
				StretchItems().
				Spacing(84).
				Content(),
		)
}
