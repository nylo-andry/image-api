package handlers

import (
	"net/http"
	"time"
)

// Logger wraps the a request handler call and logs the exec time once it is completed.
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		h.logger.Printf("Handled request in %s\n", time.Now().Sub(startTime))
	}
}
