package imageapi

import (
	"image"
)

type ImageService interface {
	Decolorize(img image.Image) *image.RGBA
	Pixelate(img image.Image) *image.RGBA
}
