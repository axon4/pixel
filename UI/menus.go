package UI

import (
	"errors"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func BuildNewMenu(application *Initialisation) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			width, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("must be a positive integer")
			}
			if width <= 0 {
				return errors.New("must be > 0")
			}
			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator
		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator
		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)
		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}

		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(OK bool) {
			if OK {
				pixelWidth := 0
				pixelHeight := 0

				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("inValid width"), application.Window)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)
				}

				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("inValid height"), application.Window)
				} else {
					pixelHeight, _ = strconv.Atoi(heightEntry.Text)
				}

				application.Canvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, application.Window)
	})
}

func BuildMenus(application *Initialisation) *fyne.Menu {
	return fyne.NewMenu("File", BuildNewMenu(application))
}

func SetUpMenus(application *Initialisation) {
	menus := BuildMenus(application)
	mainMenu := fyne.NewMainMenu(menus)
	application.Window.SetMainMenu(mainMenu)
}