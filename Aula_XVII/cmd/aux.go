package cmd

import "github.com/spf13/cobra"

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
