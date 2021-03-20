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
)

var (
	file      string
	hasHeader bool
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
		return checkCSVFile(fileName, hasH)
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

func checkCSVFile(f string, hasH bool) error {
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
		if _, err := dataparser.Read(product.Image); err != nil { //TODO: don't forget to concatenate images path with image name!
			fmt.Printf(
				"Image of row %d couldn't be found or open\n",
				index+adder,
			)
		}
	}
	return nil
}
