package UI

import (
	"image/color"
	"pixel/swatch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func BuildSwatches(application *Initialisation) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)

	for i := 0; i < cap(application.Swatches); i++ {
		initialColour := color.NRGBA{255, 255, 255, 255}
		newSwatch := swatch.NewSwatch(application.State, i, initialColour, func(clickedSwatch *swatch.Swatch) {
			for j := 0; j < len(application.Swatches); j++ {
				application.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}

			application.State.SwatchSelected = clickedSwatch.SwatchIndex
			application.State.BrushColour = clickedSwatch.Colour
		})

		if i == 0 {
			newSwatch.Selected = true
			application.State.SwatchSelected = 0
			newSwatch.Refresh()
		}

		application.Swatches = append(application.Swatches, newSwatch)
		canvasSwatches = append(canvasSwatches, newSwatch)
	}

	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}