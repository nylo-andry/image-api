package images

import (
	"image"
	"image/color"
	"math"
	"sync"
)

const chunkCount = 50

// Decolorize returns a greyscaled version of an the image received in the params.
// The only supported image format is JPG.
func (i *ImageService) Decolorize(img image.Image) *image.RGBA {
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

	return greyscaledImage
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
