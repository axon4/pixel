package UI

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lusingander/colorpicker"
)

func SetUpColourPicker(application *Initialisation) *fyne.Container {
	picker := colorpicker.New(200, colorpicker.StyleHue)
	picker.SetOnChanged(func(colour color.Color) {
		application.State.BrushColour = colour
		application.Swatches[application.State.SwatchSelected].SetColour(colour)
	})

	return container.NewVBox(picker)
}