package main

import (
	"image/color"
	"pixel/UI"
	"pixel/configuration"
	"pixel/swatch"

	"fyne.io/fyne/v2/app"
)

func main() {
	application := app.New()
	window := application.NewWindow("pixel")
	state := configuration.State{
		BrushColour:    color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}
	initialisation := UI.Initialisation{
		Window:   window,
		State:    &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	UI.SetUp(&initialisation)
	initialisation.Window.ShowAndRun()
}