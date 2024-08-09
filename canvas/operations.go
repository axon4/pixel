package canvas

import "fyne.io/fyne/v2"

func (canvas *Canvas) Pan(previousCoOrdinate, currentCoOrdinate fyne.PointEvent) {
	deltaX := currentCoOrdinate.Position.X - previousCoOrdinate.Position.X
	deltaY := currentCoOrdinate.Position.Y - previousCoOrdinate.Position.Y

	canvas.CanvasOffSet.X += deltaX
	canvas.CanvasOffSet.Y += deltaY
	canvas.Refresh()
}

func (canvas *Canvas) zoom(direction int) {
	switch {
	case direction > 0:
		canvas.PixelSize += 1
	case direction < 0:
		if canvas.PixelSize > 2 {
			canvas.PixelSize -= 1
		}
	default:
		canvas.PixelSize = 10
	}
}