package ui

import (
	"fyne.io/fyne/v2"
	"github.com/zohaibAsif/pixl/apptype"
	"github.com/zohaibAsif/pixl/pxcanvas"
	"github.com/zohaibAsif/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
