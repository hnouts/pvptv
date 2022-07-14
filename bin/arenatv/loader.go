package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type loader struct {
	app.Compo

	Iloading     bool
	Iclass       string
	Ititle       string
	Idescription string
	Ierr         error
	Isize        int
}

func newLoader() *loader {
	return &loader{
		Isize: 66,
	}
}

func (l *loader) Class(v string) *loader {
	if v == "" {
		return l
	}
	if l.Iclass != "" {
		l.Iclass += " "
	}
	l.Iclass += v
	return l
}

func (l *loader) Title(v string) *loader {
	l.Ititle = v
	return l
}

func (l *loader) Description(v string) *loader {
	l.Idescription = v
	return l
}

func (l *loader) Loading(v bool) *loader {
	l.Iloading = v
	return l
}

func (l *loader) Err(err error) *loader {
	l.Ierr = err
	return l
}

func (l *loader) Size(v int) *loader {
	l.Isize = v
	return l
}

func (l *loader) Render() app.UI {
	display := "hide"
	if l.Iloading {
		display = ""
	}

	title := l.Ititle
	descriptionMsg := l.Idescription
	descriptionColor := ""
	state := "active"

	if l.Ierr != nil {
		title += " failed"
		descriptionMsg = l.Ierr.Error()
		descriptionColor = "error"
		state = "inactive"
		display = ""
	}

	return app.Aside().
		Class("loader").
		Class(display).
		Class(l.Iclass).
		Body(
			ui.Stack().
				Class("fill").
				Center().
				Middle().
				Content(
					app.Div().
						Style("width", fmt.Sprintf("%vpx", l.Isize)).
						Style("height", fmt.Sprintf("%vpx", l.Isize)).
						Body(
							app.Div().
								Class("icon").
								Class(state).
								Style("width", fmt.Sprintf("%vpx", l.Isize-2)).
								Style("height", fmt.Sprintf("%vpx", l.Isize-2)),
						),
					app.Div().
						Class("hspace-in").
						Body(
							app.Header().
								Class("h2").
								Text(title),
							app.P().
								Class(descriptionColor).
								Text(descriptionMsg),
						),
				),
		)
}
