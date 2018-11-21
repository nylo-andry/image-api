package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// NewfileUploadRequest creates a client request to the api of the project
// with a predetermined file.
func NewfileUploadRequest(filePath string) (*http.Request, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", ".")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "/images", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
