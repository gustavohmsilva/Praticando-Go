package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XIV/client/config"
	"github.com/gustavohmsilva/Praticando-Go/Aula_XIV/client/dataparser"
	"github.com/gustavohmsilva/Praticando-Go/Aula_XIV/client/uploader"
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
	status := response.StatusCode
	fmt.Println(status, response.Body)
	path := fmt.Sprintf("%s/", conf.GetString("images_path"))
	for _, p := range ps {
		fileToUpload, err := uploader.NewUplRequestSrc(
			conf.GetString("image_endpoint"),
			"file",
			path,
			p.Image,
		)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		req, err := uploader.NewFileReq(fileToUpload)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		fmt.Println(resp.StatusCode)
	}
	if err != nil {
		panic(err)
	}
}
