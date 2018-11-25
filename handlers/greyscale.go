package handlers

import (
	"image"
	"image/jpeg"
	"net/http"
)

// Greyscale handles the request to remove all colors of an image.
func (h *Handlers) Greyscale(w http.ResponseWriter, r *http.Request) {
	i, ok := r.Context().Value(imageContextKey).(image.Image)
	if !ok {
		h.logger.Println("Received a non-image parameter from request context")
		http.Error(w, "An error occured while reading the image.", http.StatusInternalServerError)
		return
	}

	img := h.ImageService.Decolorize(i)

	jpeg.Encode(w, img, nil)
}
