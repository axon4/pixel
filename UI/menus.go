package UI

import (
	"errors"
	"image"
	"image/png"
	"os"
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

func saveFileDialogue(application *Initialisation) {
	dialog.ShowFileSave(func(URI fyne.URIWriteCloser, err error) {
		if URI == nil {
			return
		} else {
			err := png.Encode(URI, application.Canvas.PixelData)

			if err != nil {
				dialog.ShowError(err, application.Window)

				return
			}

			application.State.SetFilePath(URI.URI().Path())
		}
	}, application.Window)
}

func BuildSaveMenu(application *Initialisation) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if application.State.FilePath == "" {
			saveFileDialogue(application)
		} else {
			tryClose := func(file *os.File) {
				err := file.Close()

				if err != nil {
					dialog.ShowError(err, application.Window)
				}
			}

			file, err := os.Create(application.State.FilePath)
			defer tryClose(file)

			if err != nil {
				dialog.ShowError(err, application.Window)

				return
			}

			err = png.Encode(file, application.Canvas.PixelData)

			if err != nil {
				dialog.ShowError(err, application.Window)

				return
			}
		}
	})
}

func BuildSaveAsMenu(application *Initialisation) *fyne.MenuItem {
	return fyne.NewMenuItem("Save As", func() {
		saveFileDialogue(application)
	})
}

func BuildOpenMenu(application *Initialisation) *fyne.MenuItem {
	return fyne.NewMenuItem("Open", func() {
		dialog.ShowFileOpen(func(URI fyne.URIReadCloser, err error) {
			if URI == nil {
				return
			} else {
				image, _, err := image.Decode(URI)

				if err != nil {
					dialog.ShowError(err, application.Window)

					return
				}

				application.Canvas.LoadImage(image)
				application.State.SetFilePath(URI.URI().Path())
				imageColours := GetImageColours(image)
				var i int

				for c := range imageColours {
					if i == len(application.Swatches) {
						break
					}

					application.Swatches[i].SetColour(c)
					i++
				}
			}
		}, application.Window)
	})
}

func BuildMenus(application *Initialisation) *fyne.Menu {
	return fyne.NewMenu("File", BuildNewMenu(application), BuildSaveMenu(application), BuildSaveAsMenu(application), BuildOpenMenu(application))
}

func SetUpMenus(application *Initialisation) {
	menus := BuildMenus(application)
	mainMenu := fyne.NewMainMenu(menus)
	application.Window.SetMainMenu(mainMenu)
}