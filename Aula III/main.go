package main

import "fmt"

type calculo struct {
	diametroTubo    float64
	compArranhador  float64
	grossCordaSisal float64
}

func calculaMetragem(m calculo) float64 {
	res := (m.compArranhador / m.grossCordaSisal) * (2 * 3.14 * (m.diametroTubo / 2))
	return float64(int(res)) / 100
}

func main() {
	var medidas calculo
	fmt.Printf("Diâmetro do Tubo Kraft: ")
	fmt.Scanf("%f\n", &medidas.diametroTubo)
	fmt.Printf("Comprimento do Arranhador: ")
	fmt.Scanf("%f\n", &medidas.compArranhador)
	fmt.Printf("Grossura da Corda de Sisal: ")
	fmt.Scanf("%f\n", &medidas.grossCordaSisal)

	resultado := calculaMetragem(medidas)
	fmt.Printf("Será necessário %.1f Metros de corda de sisal por arranhador", resultado)
}
