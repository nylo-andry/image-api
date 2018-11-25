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

func Test_NoImage(t *testing.T) {
	h := NewHandlers(logger, imageService)
	ts := httptest.NewServer(h.MultipartFormMiddleware(h.ImageRequestMiddleware(testHandler())))
	defer ts.Close()

	fakeContentType := "multipart/form-data; boundary=--"
	resp, err := http.Post(ts.URL, fakeContentType, nil)
	if err != nil {
		panic(err)
	}

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func testHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func Test_UnsupportedFormat(t *testing.T) {
	h := NewHandlers(logger, imageService)
	ts := httptest.NewServer(h.MultipartFormMiddleware(h.ImageRequestMiddleware(testHandler())))
	defer ts.Close()

	filePath, err := getFileToUploadPath("blank.png")
	if err != nil {
		t.Fatal(err)
	}

	body, contentType, err := test.NewFileUploadRequest(filePath)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(ts.URL, contentType, body)
	if err != nil {
		panic(err)
	}

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestImages_AllOk(t *testing.T) {
	h := NewHandlers(logger, imageService)
	ts := httptest.NewServer(h.MultipartFormMiddleware(h.ImageRequestMiddleware(testHandler())))
	defer ts.Close()

	filePath, err := getFileToUploadPath("blank.jpg")
	if err != nil {
		t.Fatal(err)
	}

	body, contentType, err := test.NewFileUploadRequest(filePath)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(ts.URL, contentType, body)
	if err != nil {
		panic(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func getFileToUploadPath(fileName string) (string, error) {
	return filepath.Abs("../test/testdata/" + fileName)
}
