package cmd

import (
	"fmt"
	"uploader/dataparser"

	"github.com/spf13/cobra"
)

const (
	// FILE is name of the file flag
	FILE = "file"
	// HASHEADER is the name of the flag for the CSV header
	HASHEADER = "header"
	// IMGPATH is the name of the flat that receives the images path
	// location
	IMGPATH = "images"
)

var (
	file      string
	hasHeader bool
	imgPath   string
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check if the csv file is valid",
	Long: `The check command basically validates if the csv files is a valid
csv file to be used by the jamjam store`,
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
		return checkCSVFile(fileName, hasH, imgsPath)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(
		&file,
		FILE,
		string(FILE[0]),
		"",
		"The csv file to be validated",
	)
	checkCmd.Flags().BoolVarP(
		&hasHeader,
		HASHEADER,
		string(HASHEADER[1]),
		false,
		"If the CSV file has a col title header or not",
	)
	checkCmd.MarkFlagRequired(FILE)
	checkCmd.Flags().StringVarP(
		&imgPath,
		IMGPATH,
		string(IMGPATH[0]),
		"",
		"Mandatory path where the images from the CSV are located",
	)
}

func getImagesPath(c *cobra.Command) (string, error) {
	imagesPath, err := c.Flags().GetString(IMGPATH)
	if err != nil {
		return "", err
	}
	return imagesPath, nil
}

func getFilename(c *cobra.Command) (string, error) {
	fileName, err := c.Flags().GetString(FILE)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func getHasHeader(c *cobra.Command) (bool, error) {
	header, err := c.Flags().GetBool(HASHEADER)
	if err != nil {
		return false, err
	}
	return header, nil
}

func checkCSVFile(f string, hasH bool, imgsPath string) error {
	file, err := dataparser.Read(f)
	if err != nil {
		return err
	}
	products, err := dataparser.CSV(file, hasH)
	if err != nil {
		return err
	}
	adder := 1
	if hasH {
		adder++
	}
	if string(imgsPath[len(imgsPath)-1]) != "/" {
		imgsPath = fmt.Sprintf("%s/", imgsPath)
	}
	for index, product := range products {
		if product.Name == "" {
			fmt.Printf(
				"Product name of row %d faulty\n",
				index+adder,
			)
		}
		if product.Desc == "" {
			fmt.Printf(
				"Product description of row %d faulty\n",
				index+adder,
			)
		}
		if product.Category == "" {
			fmt.Printf(
				"Product category of row %d doesn't exist\n",
				index+adder,
			)
		}
		if product.Image == "" {
			fmt.Printf(
				"Product image of row %d can't be empty\n",
				index+adder,
			)
		}
		p := fmt.Sprintf("%s%s", imgsPath, product.Image)
		if _, err := dataparser.Read(p); err != nil {
			fmt.Printf(
				"Image of row %d couldn't be found or open\n",
				index+adder,
			)
		}
	}
	return nil
}
