package canvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type CanvasRenderer struct {
	canvas       *Canvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject
}

func (renderer *CanvasRenderer) MinSize() fyne.Size {
	return renderer.canvas.DrawingArea
}

func (renderer *CanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)

	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}

	objects = append(objects, renderer.canvasImage)
	objects = append(objects, renderer.canvasCursor...)

	return objects
}

func (renderer *CanvasRenderer) LayOutCanvas(size fyne.Size) {
	imagePixelWidth := renderer.canvas.PixelColumns
	imagePixelHeight := renderer.canvas.PixelRows
	pixelSize := renderer.canvas.PixelSize

	renderer.canvasImage.Move(fyne.NewPos(renderer.canvas.CanvasOffSet.X, renderer.canvas.CanvasOffSet.Y))
	renderer.canvasImage.Resize(fyne.NewSize(float32(imagePixelWidth*pixelSize), float32(imagePixelHeight*pixelSize)))
}

func (renderer *CanvasRenderer) LayOutBorder(size fyne.Size) {
	offSet := renderer.canvas.CanvasOffSet
	imageWidth := renderer.canvasImage.Size().Width
	imageHeight := renderer.canvasImage.Size().Height

	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offSet.X, offSet.Y)
	left.Position2 = fyne.NewPos(offSet.X, offSet.Y+imageHeight)

	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offSet.X, offSet.Y)
	top.Position2 = fyne.NewPos(offSet.X+imageWidth, offSet.Y)

	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offSet.X+imageWidth, offSet.Y)
	right.Position2 = fyne.NewPos(offSet.X+imageWidth, offSet.Y+imageHeight)

	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offSet.X, offSet.Y+imageHeight)
	bottom.Position2 = fyne.NewPos(offSet.X+imageWidth, offSet.Y+imageHeight)
}

func (renderer *CanvasRenderer) Layout(size fyne.Size) {
	renderer.LayOutCanvas(size)
	renderer.LayOutBorder(size)
}

func (renderer *CanvasRenderer) Refresh() {
	if renderer.canvas.reLoadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.canvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels
		renderer.canvasImage.FillMode = canvas.ImageFillContain
		renderer.canvas.reLoadImage = false
	}

	renderer.Layout(renderer.canvas.Size())
	canvas.Refresh(renderer.canvasImage)
}

func (renderer *CanvasRenderer) Destroy() {}