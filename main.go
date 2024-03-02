package main

import (
	"log"
	"os"

	"toucam/windows"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

var ops op.Ops

func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Egg Timer"),
			//app.Maximized.Option(),
		)

		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	theme := material.NewTheme()

	mw := windows.NewMainWindow(theme)

	for {
		switch e := w.NextEvent().(type) {
		case app.FrameEvent:
			mw.Layout(app.NewContext(&ops, e))
		case app.DestroyEvent:
			return e.Err
		}
	}
}
