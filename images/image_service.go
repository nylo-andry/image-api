package images

import imageapi "github.com/nylo-andry/image-api"

var _ imageapi.ImageService = &ImageService{}

type ImageService struct{}
