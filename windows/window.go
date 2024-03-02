package windows

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type window struct {
	title  string
	theme  *material.Theme
	Layout func(gtx layout.Context) layout.Dimensions
}
