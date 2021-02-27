package uploader

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XVI/client/dataparser"
)

// UplRequestSrc define the necessary data for a file upload http request to be
// created.
type UplRequestSrc struct {
	Endpoint      string
	FileParameter string
	Path          string
	FileName      string
	Status        error
}

// NewUplRequestSrc create a populated UplRequestSrc object
func NewUplRequestSrc(
	uri, fileParam, path, fileName string,
) (
	UplRequestSrc, error,
) {
	if _, err := url.ParseRequestURI(uri); err != nil {
		return UplRequestSrc{}, err
	}
	f, err := dataparser.Read(fmt.Sprintf("%s/%s", path, fileName))
	f.Close()
	if err != nil {
		return UplRequestSrc{}, err
	}
	if fileParam == "" {
		return UplRequestSrc{}, fmt.Errorf("fileParam is mandatory")
	}
	return UplRequestSrc{uri, fileParam, path, fileName, nil}, nil
}

// NewFileReq creates a http request with the necessary data to upload a file to
// a server over http
func NewFileReq(source UplRequestSrc) (*http.Request, error) {
	f, err := os.Open(fmt.Sprintf("%s/%s", source.Path, source.FileName))
	if err != nil {
		return &http.Request{}, err
	}
	defer f.Close()
	descStruct, err := f.Stat()
	if err != nil {
		return &http.Request{}, err
	}
	reqBody := new(bytes.Buffer)
	bufferWriter := multipart.NewWriter(reqBody)
	formData, err := bufferWriter.CreateFormFile(
		source.FileParameter,
		descStruct.Name(),
	)
	if err != nil {
		return &http.Request{}, err
	}
	io.Copy(formData, f)
	err = bufferWriter.Close()
	if err != nil {
		return &http.Request{}, err
	}
	request, err := http.NewRequest("POST", source.Endpoint, reqBody)
	request.Header.Set("Content-Type", bufferWriter.FormDataContentType())
	return request, nil
}
