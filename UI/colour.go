package UI

import (
	"image"
	"image/color"
)

func GetImageColours(image image.Image) map[color.Color]struct{} {
	colours := make(map[color.Color]struct{})
	var empty struct{}

	for y := 0; y < image.Bounds().Dy(); y++ {
		for x := 0; x < image.Bounds().Dx(); x++ {
			colours[image.At(x, y)] = empty
		}
	}

	return colours
}