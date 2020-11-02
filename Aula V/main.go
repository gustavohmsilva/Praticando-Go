package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputParameters := os.Args[1:]
	processedData, err := parseData(inputParameters)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	median, _ := calculateMedian(processedData)
	fmt.Println(median)
}

func calculateMedian(dat []float64) (float64, error) {
	var sum float64
	for _, f := range dat {
		sum += f
	}
	div := float64(len(dat))
	return sum / div, nil
}

func parseData(input []string) ([]float64, error) {
	var pd []float64
	for _, v := range input {
		v := strings.ReplaceAll(v, ",", ".")
		pv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return pd, fmt.Errorf("Incorrect Data: failed processing value %v", v)
		}
		pd = append(pd, pv)
	}
	return pd, nil
}
