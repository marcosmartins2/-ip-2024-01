package main

import "fmt"

func main() {
	var resultado float64
	coeficiente := 1.0
	x := obterX()
	N := obterN()

	for n := 0; n <= N; n++ {
		resultado += coeficiente / float64(calcularFatorial(n))
		coeficiente *= x
	}

	fmt.Printf("e^%.2f = %.6f\n", x, resultado)
}

func obterX() float64 {
	var valorX float64
	fmt.Println("Escolha um valor positivo para x:")
	fmt.Scanln(&valorX)
	if valorX >= 0 {
		return valorX
	}
	fmt.Println("O valor deve ser positivo. Tente novamente.")
	return obterX()
}

func obterN() int {
	var N int
	fmt.Println("Escolha um valor inteiro positivo para N:")
	fmt.Scanln(&N)
	if N >= 0 {
		return N
	}
	fmt.Println("O valor deve ser inteiro positivo. Tente novamente.")
	return obterN()
}

func calcularFatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calcularFatorial(n-1)
}
