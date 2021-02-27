package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XVI/client/config"
	"github.com/gustavohmsilva/Praticando-Go/Aula_XVI/client/dataparser"
	"github.com/gustavohmsilva/Praticando-Go/Aula_XVI/client/uploader"
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
	if err != nil {
		panic(err)
	}
	status := response.StatusCode
	fmt.Println(status, response.Body)
	path := fmt.Sprintf("%s", conf.GetString("images_path"))
	wg := sync.WaitGroup{}
	wg.Add(len(ps))
	ch := make(chan *uploader.UplRequestSrc, len(ps))
	for _, p := range ps {
		go func(ep, p, i string, w *sync.WaitGroup, c chan *uploader.UplRequestSrc) {
			defer w.Done()
			fileToUpload, err := uploader.NewUplRequestSrc(ep, "file", p, i)
			if err != nil {
				fileToUpload.FileName = i
				fileToUpload.Status = err
				c <- &fileToUpload
				return
			}
			req, err := uploader.NewFileReq(fileToUpload)
			if err != nil {
				fileToUpload.Status = err
				c <- &fileToUpload
				return
			}
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				fileToUpload.Status = err
				c <- &fileToUpload
				return
			}
			c <- &fileToUpload
		}(conf.GetString("image_endpoint"), path, p.Image, &wg, ch)
	}
	wg.Wait()
	close(ch)
	for v := range ch {
		if v.Status != nil {
			fmt.Printf("%s failed... Reason: %v\n", v.FileName, v.Status.Error())
			continue
		}
		fmt.Printf("%s send successfully!\n", v.FileName)
	}
}
