package main

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/product", saveProduct)
	e.Logger.Fatal(e.Start(":8080"))
}

func saveProduct(cnx echo.Context) error {
	type product struct {
		Name     string `json:"name"`
		Category string `json:"category"`
		Desc     string `json:"desc"`
	}
	var ps []product
	err := cnx.Bind(&ps)
	if err != nil {
		return err
	}
	fmt.Println(ps)
	return nil
}
