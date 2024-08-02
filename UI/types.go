package UI

import (
	"pixel/configuration"
	"pixel/swatch"

	"fyne.io/fyne/v2"
)

type Initialisation struct {
	Window   fyne.Window
	State    *configuration.State
	Swatches []*swatch.Swatch
}