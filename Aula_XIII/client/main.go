package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XI/dataparser"
	"github.com/gustavohmsilva/Praticando-Go/Aula_XIII/client/config"
	"github.com/spf13/viper"
)

func main() {
	config.Load()
	f, err := dataparser.Read(viper.GetString("filename"))
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
		viper.GetString("server"),
		"application/json",
		bytes.NewBuffer(jsonPs),
	)
	if err != nil {
		panic(err)
	}
	status := response.StatusCode
	fmt.Println(status, response.Body)
}
