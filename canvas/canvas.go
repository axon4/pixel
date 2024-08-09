package canvas

import (
	"image"
	"image/color"
	"pixel/configuration"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type CanvasMouseState struct {
	previousCoOrdinate *fyne.PointEvent
}

type Canvas struct {
	widget.BaseWidget
	configuration.CanvasConfiguration
	renderer         *CanvasRenderer
	PixelData        image.Image
	mouseState       CanvasMouseState
	applicationState *configuration.State
	reLoadImage      bool
	showMouse        bool
}

func (canvas *Canvas) Bounds() image.Rectangle {
	x0 := int(canvas.CanvasOffSet.X)
	y0 := int(canvas.CanvasOffSet.Y)
	x1 := int(canvas.PixelColumns*canvas.PixelSize + int(canvas.CanvasOffSet.X))
	y1 := int(canvas.PixelRows*canvas.PixelSize + int(canvas.CanvasOffSet.Y))

	return image.Rect(x0, y0, x1, y1)
}

func InBounds(position fyne.Position, bounds image.Rectangle) bool {
	if position.X >= float32(bounds.Min.X) && position.X < float32(bounds.Max.X) && position.Y >= float32(bounds.Min.Y) && position.Y < float32(bounds.Max.Y) {
		return true
	} else {
		return false
	}
}

func NewBlankImage(columns, rows int, colour color.Color) image.Image {
	blankImage := image.NewNRGBA(image.Rect(0, 0, columns, rows))

	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			blankImage.Set(x, y, colour)
		}
	}

	return blankImage
}

func NewCanvas(state *configuration.State, configuration configuration.CanvasConfiguration) *Canvas {
	canvas := &Canvas{
		CanvasConfiguration: configuration,
		applicationState:    state,
	}
	canvas.PixelData = NewBlankImage(canvas.PixelColumns, canvas.PixelRows, color.NRGBA{128, 128, 128, 255})
	canvas.ExtendBaseWidget(canvas)

	return canvas
}

func (pixelCanvas *Canvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pixelCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain
	canvasBorder := make([]canvas.Line, 4)

	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &CanvasRenderer{
		canvas:       pixelCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}
	pixelCanvas.renderer = renderer

	return renderer
}

func (canvas *Canvas) TryPan(previousCoOrdinate *fyne.PointEvent, event *desktop.MouseEvent) {
	if previousCoOrdinate != nil && event.Button == desktop.MouseButtonTertiary {
		canvas.Pan(*previousCoOrdinate, event.PointEvent)
	}
}

func (canvas *Canvas) SetColour(colour color.Color, x, y int) {
	if NRGBA, OK := canvas.PixelData.(*image.NRGBA); OK {
		NRGBA.Set(x, y, colour)
	}

	if RGBA, OK := canvas.PixelData.(*image.RGBA); OK {
		RGBA.Set(x, y, colour)
	}

	canvas.Refresh()
}

func (canvas *Canvas) MouseToCanvasXY(event *desktop.MouseEvent) (*int, *int) {
	bounds := canvas.Bounds()

	if !InBounds(event.Position, bounds) {
		return nil, nil
	}

	pixelSize := float32(canvas.PixelSize)
	offSetX := canvas.CanvasOffSet.X
	offSetY := canvas.CanvasOffSet.Y
	x := int((event.Position.X - offSetX) / pixelSize)
	y := int((event.Position.Y - offSetY) / pixelSize)

	return &x, &y
}