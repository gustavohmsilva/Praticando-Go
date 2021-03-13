package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	// FILE is name of the file flag
	FILE = "file"
)

var file string

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check if the csv file is valid",
	Long: `The check command basically validates if the csv files is a valid
csv file to be used by the jamjam store`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName, err := cmd.Flags().GetString(FILE)
		if err != nil {
			return err
		}
		fmt.Println(fileName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(
		&file,
		FILE,
		"f",
		"",
		"The csv file to be validated",
	)
	checkCmd.MarkFlagRequired(FILE)
}
