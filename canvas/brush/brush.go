package brush

import (
	"pixel/configuration"

	"fyne.io/fyne/v2/driver/desktop"
)

const Pixel = iota

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