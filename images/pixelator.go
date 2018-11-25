package images

import (
	"errors"
	"image"
	"io"

	"golang.org/x/image/draw"
)

func (i *ImageService) Pixelate(file io.Reader) (*image.RGBA, error) {
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	if format != "jpeg" {
		return nil, errors.New("Only JPG files are supported")
	}

	b := img.Bounds()
	resized := image.NewRGBA(image.Rect(0, 0, b.Dx()/4, b.Dy()/4))
	draw.NearestNeighbor.Scale(resized, resized.Bounds(), img, b, draw.Over, nil)

	pixelated := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.NearestNeighbor.Scale(pixelated, pixelated.Bounds(), resized, resized.Bounds(), draw.Over, nil)

	return pixelated, nil
}
