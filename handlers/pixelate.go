package handlers

import (
	"image/jpeg"
	"net/http"
)

func (h *Handlers) Pixelate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(maxImageSize)
	file, _, err := r.FormFile("image")
	if err != nil {
		h.logger.Printf("An error occured while reading the file: %v", err)
		http.Error(w, "Could not read file.", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	img, err := h.ImageService.Pixelate(file)
	if err != nil {
		h.logger.Printf("An error occured while decolorizing the file: %v", err)
		http.Error(w, "Unsupported file format. Only JPEG files are allowed.", http.StatusBadRequest)
		return
	}

	jpeg.Encode(w, img, nil)
}
