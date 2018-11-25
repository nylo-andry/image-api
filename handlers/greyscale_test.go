package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/nylo-andry/image-api/images"
	"github.com/nylo-andry/image-api/test"
)

var logger = log.New(os.Stdout, "image-api::", log.Lshortfile)
var imageService = &images.ImageService{}

func TestGreyscale_NoImage(t *testing.T) {
	h := NewHandlers(logger, imageService)
	req, err := http.NewRequest("POST", "/images", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Greyscale)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestImages_UnsupportedFormat(t *testing.T) {
	h := NewHandlers(logger, imageService)
	filePath, err := getFileToUploadPath("blank.png")
	if err != nil {
		t.Fatal(err)
	}

	req, err := test.NewFileUploadRequest(filePath)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Greyscale)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestImages_AllOk(t *testing.T) {
	h := NewHandlers(logger, imageService)
	filePath, err := getFileToUploadPath("blank.jpg")
	if err != nil {
		t.Fatal(err)
	}

	req, err := test.NewFileUploadRequest(filePath)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.Greyscale)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func getFileToUploadPath(fileName string) (string, error) {
	return filepath.Abs("../test/testdata/" + fileName)
}
