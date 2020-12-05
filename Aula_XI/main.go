package main

import (
	"fmt"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XI/database"
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
	for _, p := range ps {
		id, err := database.InsertProduct(p)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Product saved in the database with the ID %d\n", id)
	}
}
