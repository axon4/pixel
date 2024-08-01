package swatch

import (
	"image/color"
	"pixel/configuration"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Swatch struct {
	widget.BaseWidget
	SwatchIndex  int
	Colour       color.Color
	Selected     bool
	clickHandler func(swatch *Swatch)
}

func (swatch *Swatch) SetColour(colour color.Color) {
	swatch.Colour = colour
	swatch.Refresh()
}

func NewSwatch(state *configuration.State, swatchIndex int, colour color.Color, clickHandler func(swatch *Swatch)) *Swatch {
	swatch := &Swatch{
		SwatchIndex:  swatchIndex,
		Colour:       colour,
		Selected:     false,
		clickHandler: clickHandler,
	}
	swatch.ExtendBaseWidget(swatch)

	return swatch
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Colour)
	objects := []fyne.CanvasObject{square}

	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}