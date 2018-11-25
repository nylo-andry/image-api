package handlers

import (
	"context"
	"net/http"
)

const (
	maxImageSize = 2 << 6
)

func (h *Handlers) MultipartFormMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r.ParseMultipartForm(maxImageSize)
		file, _, err := r.FormFile("image")
		if err != nil {
			h.logger.Printf("An error occured while reading the file: %v", err)
			http.Error(w, "Could not read file.", http.StatusBadRequest)
			return
		}
		defer file.Close()

		ctx = context.WithValue(ctx, fileContextKey, file)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
