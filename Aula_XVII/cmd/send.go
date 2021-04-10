package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"uploader/dataparser"

	"github.com/gustavohmsilva/Praticando-Go/Aula_XVII/client/config"
	"github.com/spf13/cobra"
)

// sendCmd represent the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "deliver products to the store's http server",
	Long: `The send command basically start delivering the products and it's
images to the jamjam store's http server for publishing.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName, err := getFilename(cmd)
		if err != nil {
			return err
		}
		hasH, err := getHasHeader(cmd)
		if err != nil {
			return err
		}
		imgsPath, err := getImagesPath(cmd)
		if err != nil {
			return err
		}

		products, err := fetchProducts(fileName, hasH)
		if err != nil {
			return err
		}

		err = sendProductInfo(products)
		if err != nil {
			return err
		}
		err = sendProductImages(products, imgsPath)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(
		&file,
		FILE,
		string(FILE[0]),
		"",
		"The csv file to be validated",
	)
	sendCmd.MarkFlagRequired(FILE)
	sendCmd.Flags().BoolVarP(
		&hasHeader,
		HASHEADER,
		string(HASHEADER[1]),
		false,
		"If the CSV file has a col title header or not",
	)
	sendCmd.Flags().StringVarP(
		&imgPath,
		IMGPATH,
		string(IMGPATH[0]),
		"",
		"Mandatory path where the images from the CSV are located",
	)
}

func fetchProducts(f string, hasH bool) ([]dataparser.Product, error) {
	file, err := dataparser.Read(f)
	if err != nil {
		return nil, err
	}

	products, err := dataparser.CSV(file, hasH)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func sendProductInfo(ps []dataparser.Product) error {
	conf, err := config.Load()
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
	return nil
}

func sendProductImages(ps []dataparser.Product, imgsPath string) error {
	// Proccess delivery of text data to store
	return nil
}
