package canvas

import (
	"pixel/canvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (canvas *Canvas) Scrolled(event *fyne.ScrollEvent) {
	canvas.zoom(int(event.Scrolled.DY))
	canvas.Refresh()
}

func (canvas *Canvas) MouseIn(event *desktop.MouseEvent) {}

func (canvas *Canvas) MouseMoved(event *desktop.MouseEvent) {
	if x, y := canvas.MouseToCanvasXY(event); x != nil && y != nil {
		brush.TryBrush(canvas.applicationState, canvas, event)
	}

	canvas.TryPan(canvas.mouseState.previousCoOrdinate, event)
	canvas.Refresh()
	canvas.mouseState.previousCoOrdinate = &event.PointEvent
}

func (canvas *Canvas) MouseOut() {}

func (canvas *Canvas) MouseUp(event *desktop.MouseEvent) {}

func (canvas *Canvas) MouseDown(event *desktop.MouseEvent) {
	brush.TryBrush(canvas.applicationState, canvas, event)
}