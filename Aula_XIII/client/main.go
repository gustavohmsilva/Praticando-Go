package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XIII/client/config"
	"github.com/gustavohmsilva/Praticando-Go/Aula_XIII/client/dataparser"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}
	f, err := dataparser.Read(conf.GetString("filename"))
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
		conf.GetString("endpoint"),
		"application/json",
		bytes.NewBuffer(jsonPs),
	)
	if err != nil {
		panic(err)
	}
	status := response.StatusCode
	fmt.Println(status, response.Body)
}
