package handlers

import (
	"context"
	"image"
	"io"
	"net/http"
)

func (h *Handlers) ImageRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		file, ok := ctx.Value(fileContextKey).(io.Reader)
		if !ok {
			h.logger.Printf("Received a non-file parameter")
			http.Error(w, "Could not read image.", http.StatusInternalServerError)
			return
		}

		img, format, err := image.Decode(file)
		if err != nil {
			h.logger.Printf("Could not read image: %v", err)
			http.Error(w, "Could not read image.", http.StatusBadRequest)
			return
		}

		if format != "jpeg" {
			h.logger.Println("Received a non-JPEG image")
			http.Error(w, "Only JPEG images are supported.", http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(ctx, imageContextKey, img)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
