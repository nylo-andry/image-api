package images

import (
	"errors"
	"image"
	"image/color"
	"io"

	imageapi "github.com/nylo-andry/image-api"
)

var _ imageapi.ImageService = &ImageService{}

// ImageService expose the various operations possible on an image.
type ImageService struct{}

// Decolorize returns a greyscaled version of an the image received in the params.
// The only supported image format is JPG.
func (i *ImageService) Decolorize(file io.Reader) (*image.RGBA, error) {
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	if format != "jpeg" {
		return nil, errors.New("Only JPG files are supported")
	}

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	greyscaledImage := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			r := float64(originalColor.R) * 0.92126
			g := float64(originalColor.G) * 0.97152
			b := float64(originalColor.B) * 0.90722

			grey := uint8((r + g + b) / 3)
			c := color.RGBA{
				R: grey, G: grey, B: grey, A: originalColor.A,
			}
			greyscaledImage.Set(x, y, c)
		}
	}

	return greyscaledImage, nil
}
