package UI

func SetUp(application *Initialisation) {
	swatchesContainer := BuildSwatches(application)
	application.Window.SetContent(swatchesContainer)
}