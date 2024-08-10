package brush

import (
	"image/color"
	"pixel/configuration"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

const Pixel = iota

func Cursor(configuration configuration.CanvasConfiguration, brush configuration.BrushType, event *desktop.MouseEvent, x, y int) []fyne.CanvasObject {
	var objects []fyne.CanvasObject

	switch {
	case brush == Pixel:
		pixelSize := float32(configuration.PixelSize)
		originX := (float32(x) * pixelSize) + configuration.CanvasOffSet.X
		originY := (float32(y) * pixelSize) + configuration.CanvasOffSet.Y
		cursorColour := color.NRGBA{80, 80, 80, 255}

		left := canvas.NewLine(cursorColour)
		left.StrokeWidth = 3
		left.Position1 = fyne.NewPos(originX, originY)
		left.Position2 = fyne.NewPos(originX, originY+pixelSize)

		top := canvas.NewLine(cursorColour)
		top.StrokeWidth = 3
		top.Position1 = fyne.NewPos(originX, originY)
		top.Position2 = fyne.NewPos(originX+pixelSize, originY)

		right := canvas.NewLine(cursorColour)
		right.StrokeWidth = 3
		right.Position1 = fyne.NewPos(originX+pixelSize, originY)
		right.Position2 = fyne.NewPos(originX+pixelSize, originY+pixelSize)

		bottom := canvas.NewLine(cursorColour)
		bottom.StrokeWidth = 3
		bottom.Position1 = fyne.NewPos(originX, originY+pixelSize)
		bottom.Position2 = fyne.NewPos(originX+pixelSize, originY+pixelSize)

		objects = append(objects, left, top, right, bottom)
	}

	return objects
}

func TryPaintPixel(state *configuration.State, canvas configuration.BrushAble, event *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(event)

	if x != nil && y != nil && event.Button == desktop.MouseButtonPrimary {
		canvas.SetColour(state.BrushColour, *x, *y)

		return true
	} else {
		return false
	}
}

func TryBrush(state *configuration.State, canvas configuration.BrushAble, event *desktop.MouseEvent) bool {
	switch {
	case state.BrushType == Pixel:
		return TryPaintPixel(state, canvas, event)
	default:
		return false
	}
}