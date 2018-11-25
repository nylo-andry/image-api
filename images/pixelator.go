package images

import (
	"image"

	"golang.org/x/image/draw"
)

func (i *ImageService) Pixelate(img image.Image) *image.RGBA {
	b := img.Bounds()
	resized := image.NewRGBA(image.Rect(0, 0, b.Dx()/4, b.Dy()/4))
	draw.NearestNeighbor.Scale(resized, resized.Bounds(), img, b, draw.Over, nil)

	pixelated := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.NearestNeighbor.Scale(pixelated, pixelated.Bounds(), resized, resized.Bounds(), draw.Over, nil)

	return pixelated
}
