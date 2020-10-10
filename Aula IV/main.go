package main

import (
	"fmt"
	"math"
)

type washers struct {
	height       float64
	extDiam      float64
	intDiam      float64
	qt           int
	freightPerKg float64
}

func shippingPrice(w washers, weight float64) string {
	extRay, intRay := w.extDiam/2, w.intDiam/2
	vol := 3.14159 * w.height * (math.Pow(extRay, 2) - math.Pow(intRay, 2))
	totalWeight := vol * float64(w.qt) * weight
	sp := totalWeight * w.freightPerKg / 1000
	return fmt.Sprintf("R$%.2f", sp)
}

func main() {
	weightByCubicCm := 7.8
	var inputData washers
	fmt.Printf("Type the washer's weight: ")
	fmt.Scanf("%f\n", &inputData.height)
	fmt.Printf("Type the washer's external diameter: ")
	fmt.Scanf("%f\n", &inputData.extDiam)
	fmt.Printf("Type the washer's internal diameter: ")
	fmt.Scanf("%f\n", &inputData.intDiam)
	fmt.Printf("Type the amount of units: ")
	fmt.Scanf("%d\n", &inputData.qt)
	fmt.Printf("Type the freight per Kilogram: ")
	fmt.Scanf("%f\n", &inputData.freightPerKg)

	fmt.Printf("Freight total Price: %s\n", shippingPrice(inputData, weightByCubicCm))
}
