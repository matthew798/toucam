package windows

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var startButton widget.Clickable

func NewMainWindow(theme *material.Theme) window {
	mainWindow := window{
		title: "TOUCAM",
		theme: theme,
		Layout: func(gtx layout.Context) layout.Dimensions {
			//Window layout goes here
			return layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &startButton, "Start")
						return btn.Layout(gtx)
					},
				),
			)
		},
	}

	return mainWindow
}
