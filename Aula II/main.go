package main

import "fmt"

type coins struct {
	Qtd1Cents  int
	Qtd2Cents  int
	Qtd5Cents  int
	Qtd10Cents int
	Qtd20Cents int
	Qtd50Cents int
	Qtd1Euros  int
	Qtd2Euros  int
}

func countCoins(c coins) (int, int) {
	var euros int
	var cents int
	cents += c.Qtd1Cents
	cents += c.Qtd2Cents * 2
	cents += c.Qtd5Cents * 5
	cents += c.Qtd10Cents * 10
	cents += c.Qtd20Cents * 20
	cents += c.Qtd50Cents * 50
	cents += c.Qtd1Euros*100 + c.Qtd2Euros*200
	euros = cents / 100
	cents = cents % 100
	return euros, cents
}

func main() {
	var amuntOfCoins coins
	fmt.Printf("Digite quantidade de moedas de 1 cêntimos: ")
	fmt.Scanf("%d\n", &amuntOfCoins.Qtd1Cents)
	fmt.Printf("Digite quantidade de moedas de 2 cêntimos: ")
	fmt.Scanf("%d\n", &amuntOfCoins.Qtd2Cents)
	fmt.Printf("Digite quantidade de moedas de 5 cêntimos: ")
	fmt.Scanf("%d\n", &amuntOfCoins.Qtd5Cents)
	fmt.Printf("Digite quantidade de moedas de 10 cêntimos: ")
	fmt.Scanf("%d\n", &amuntOfCoins.Qtd10Cents)
	fmt.Printf("Digite quantidade de moedas de 20 cêntimos: ")
	fmt.Scanf("%d\n", &amuntOfCoins.Qtd20Cents)
	fmt.Printf("Digite quantidade de moedas de 50 cêntimos: ")
	fmt.Scanf("%d\n", &amuntOfCoins.Qtd50Cents)
	fmt.Printf("Digite quantidade de moedas de 1 euros: ")
	fmt.Scanf("%d\n", &amuntOfCoins.Qtd1Euros)
	fmt.Printf("Digite quantidade de moedas de 2 euros: ")
	fmt.Scanf("%d\n", &amuntOfCoins.Qtd2Euros)
	eur, cent := countCoins(amuntOfCoins)
	fmt.Printf("O Valor total é €%d,%d\n", eur, cent)
}
