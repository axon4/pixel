package UI

import (
	"pixel/canvas"
	"pixel/configuration"
	"pixel/swatch"

	"fyne.io/fyne/v2"
)

type Initialisation struct {
	Window   fyne.Window
	Canvas   *canvas.Canvas
	State    *configuration.State
	Swatches []*swatch.Swatch
}