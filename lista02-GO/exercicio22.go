package main

import "fmt"

func main() {
	var entrada float64
	fmt.Println("Digite o número N:")
	fmt.Scanln(&entrada)

	numerador := int(entrada * 100)
	denominador := 100

	numerador, denominador = simplificarFracao(numerador, denominador)

	fmt.Printf("%d / %d", numerador, denominador)
}

func simplificarFracao(numerador, denominador int) (int, int) {
	mdc := calcularMDC(numerador, denominador)

	numerador /= mdc
	denominador /= mdc

	return numerador, denominador
}

func calcularMDC(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
