package main

import (
	"image/color"
	"pixel/UI"
	"pixel/canvas"
	"pixel/configuration"
	"pixel/swatch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	application := app.New()
	window := application.NewWindow("pixel")
	state := configuration.State{
		BrushColour:    color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}
	canvasConfiguration := configuration.CanvasConfiguration{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffSet: fyne.NewPos(0, 0),
		PixelRows:    10,
		PixelColumns: 10,
		PixelSize:    30,
	}
	pixelCanvas := canvas.NewCanvas(&state, canvasConfiguration)
	initialisation := UI.Initialisation{
		Window:   window,
		Canvas:   pixelCanvas,
		State:    &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	UI.SetUp(&initialisation)
	initialisation.Window.ShowAndRun()
}