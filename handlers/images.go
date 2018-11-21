package handlers

import (
	"image/jpeg"
	"net/http"
)

// MaxImageSize represents the maximum allowed size of a an images which is 10 MB.
const MaxImageSize = 1 << 7

// Images handles the images modification requests and returns the greyscaled image.
func (h *Handlers) Images(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(MaxImageSize)
	file, _, err := r.FormFile("image")
	if err != nil {
		h.logger.Printf("An error occured while reading the file: %v", err)
		http.Error(w, "Could not read file.", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	img, err := h.ImageService.Decolorize(file)
	if err != nil {
		h.logger.Printf("An error occured while decolorizing the file: %v", err)
		http.Error(w, "Unsupported file format. Only JPEG files are allowed.", http.StatusBadRequest)
		return
	}

	jpeg.Encode(w, img, nil)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
