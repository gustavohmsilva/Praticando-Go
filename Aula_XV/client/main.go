package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XIV/client/config"
	"github.com/gustavohmsilva/Praticando-Go/Aula_XIV/client/dataparser"
	"github.com/gustavohmsilva/Praticando-Go/Aula_XV/client/uploader"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}
	maxProcs, err := strconv.Atoi(conf.GetString("max_procs"))
	if err != nil {
		panic(err)
	}
	runtime.GOMAXPROCS(maxProcs)
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
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(len(ps))
	for _, p := range ps {
		go func(ep, p, i string, w *sync.WaitGroup) {
			defer w.Done()
			fileToUpload, err := uploader.NewUplRequestSrc(ep, "file", p, i)
			if err != nil {
				fmt.Printf(err.Error())
			}
			req, err := uploader.NewFileReq(fileToUpload)
			if err != nil {
				fmt.Printf(err.Error())
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf(err.Error())
			}
			fmt.Println(res.StatusCode, res.Body)
		}(conf.GetString("image_endpoint"), path, p.Image, &wg)
	}
	wg.Wait()
	fmt.Println(time.Since(startTime))
}
