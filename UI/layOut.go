package UI

import "fyne.io/fyne/v2/container"

func SetUp(application *Initialisation) {
	swatchesContainer := BuildSwatches(application)
	colourPicker := SetUpColourPicker(application)
	layOut := container.NewBorder(nil, swatchesContainer, nil, colourPicker, application.Canvas)
	application.Window.SetContent(layOut)
}