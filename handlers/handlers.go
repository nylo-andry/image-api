package handlers

import "log"
import imageapi "github.com/nylo-andry/image-api"

const maxImageSize = 1 << 7

// Handlers is the representation of request handlers used in the service.
type Handlers struct {
	logger       *log.Logger
	ImageService imageapi.ImageService
}

// NewHandlers returns a new Handlers instance.
func NewHandlers(logger *log.Logger, i imageapi.ImageService) *Handlers {
	return &Handlers{logger, i}
}
