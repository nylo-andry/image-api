package imageapi

import (
	"image"
	"io"
)

type ImageService interface {
	Decolorize(file io.Reader) (*image.RGBA, error)
}