package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/product", saveProduct)
	e.POST("/image", saveImage)
	e.Logger.Fatal(e.Start(":8080"))
}

func saveProduct(cnx echo.Context) error {
	type product struct {
		Name     string `json:"name"`
		Category string `json:"category"`
		Desc     string `json:"desc"`
		Image    string `json:"image"`
	}
	var ps []product
	err := cnx.Bind(&ps)
	if err != nil {
		return err
	}
	fmt.Println(ps)
	return cnx.NoContent(http.StatusCreated)
}

func saveImage(cnx echo.Context) error {
	f, err := cnx.FormFile("file")
	if err != nil {
		return err
	}
	receivedImage, err := f.Open()
	if err != nil {
		return err
	}
	defer receivedImage.Close()
	savedImage, err := os.Create(f.Filename)
	if err != nil {
		return err
	}
	defer savedImage.Close()
	if _, err = io.Copy(savedImage, receivedImage); err != nil {
		return err
	}
	return cnx.NoContent(http.StatusCreated)
}
