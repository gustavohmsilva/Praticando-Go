package main

import "fmt"

func milesToKm(miles float64) float64 {
	return miles * 1.609
}

func main() {
	var miles float64

	fmt.Println("Converte milhas para quilômetros")
	fmt.Print("Digite o valor em milhas: ")
	fmt.Scan(&miles)
	fmt.Printf("Valor em quilômetros: %g\n", milesToKm(miles))
}
