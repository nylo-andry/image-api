package handlers

import (
	"net/http"
	"time"
)

func (h *Handlers) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		h.logger.Printf("Handled request in %s\n", time.Now().Sub(startTime))
	})
}
