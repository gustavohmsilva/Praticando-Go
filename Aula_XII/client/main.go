package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XI/dataparser"
)

func main() {
	f, err := dataparser.Read("sample_test_product.csv")
	if err != nil {
		panic(err)
	}
	ps, err := dataparser.CSV(f, true)
	if err != nil {
		panic(err)
	}
	jsonPs, err := json.Marshal(ps)
	if err != nil {
		panic(err)
	}
	response, err := http.Post(
		"http://localhost:8080/product",
		"application/json",
		bytes.NewBuffer(jsonPs),
	)
	if err != nil {
		panic(err)
	}
	status := response.StatusCode
	fmt.Println(status, response.Body)
}
