package images

import (
	"errors"
	"image"
	"image/color"
	"io"
	"math"
	"sync"

	imageapi "github.com/nylo-andry/image-api"
)

var _ imageapi.ImageService = &ImageService{}

const chunkCount = 50

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

	var wg sync.WaitGroup
	wg.Add(chunkCount)

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	chunkSize := int(math.Ceil(float64(size.Y / chunkCount)))
	greyscaledImage := image.NewRGBA(rect)

	for i := 0; i < chunkCount; i++ {
		start := i * chunkSize
		end := start + chunkSize

		go func() {
			processChunk(img, greyscaledImage, start, end)
			wg.Done()
		}()
	}

	wg.Wait()

	return greyscaledImage, nil
}

func processChunk(img image.Image, dest *image.RGBA, yStart int, yEnd int) {
	sizeX := img.Bounds().Size().X

	for x := 0; x < sizeX; x++ {
		for y := yStart; y < yEnd; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			r := float64(originalColor.R) * 0.92126
			g := float64(originalColor.G) * 0.97152
			b := float64(originalColor.B) * 0.90722

			grey := uint8((r + g + b) / 3)
			c := color.RGBA{
				R: grey, G: grey, B: grey, A: originalColor.A,
			}
			dest.Set(x, y, c)
		}
	}
}
