package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nylo-andry/image-api/handlers"
	"github.com/nylo-andry/image-api/images"
)

func main() {
	logger := log.New(os.Stdout, "image-api::", log.LstdFlags|log.Lshortfile)
	h := handlers.NewHandlers(logger, &images.ImageService{})
	r := mux.NewRouter()

	r.HandleFunc("/images/greyscale", h.Greyscale).Methods("POST")
	r.HandleFunc("/images/pixelate", h.Pixelate).Methods("POST")
	r.Use(h.Logger, h.MultipartFormMiddleware, h.ImageRequestMiddleware)

	logger.Println("server starting")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		logger.Fatalf("could not start server: %v", err)
	}
}
