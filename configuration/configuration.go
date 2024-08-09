package configuration

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type BrushType = int

type CanvasConfiguration struct {
	DrawingArea             fyne.Size
	CanvasOffSet            fyne.Position
	PixelRows, PixelColumns int
	PixelSize               int
}

type State struct {
	BrushColour    color.Color
	BrushType      int
	SwatchSelected int
	FilePath       string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}

type BrushAble interface {
	SetColour(colour color.Color, x, y int)
	MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int)
}