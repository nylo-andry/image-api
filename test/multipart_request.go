package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

// NewFileUploadRequest creates a client request to the api of the project
// with a predetermined file.
func NewFileUploadRequest(filePath string) (io.Reader, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", ".")
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}
