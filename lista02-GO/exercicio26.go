package main

import "fmt"

func main() {
	var resultadoSeno float64
	sinal := 1.0
	angulo := obterAngulo()
	coeficiente := angulo
	N := obterN()

	for n := 0; n <= N; n++ {
		resultadoSeno += sinal * coeficiente / float64(calcularFatorial(2*n+1))
		sinal *= -1
		coeficiente *= angulo * angulo
	}

	fmt.Printf("seno(%.2f) = %.6f\n", angulo, resultadoSeno)
}

func obterAngulo() float64 {
	var angulo float64
	fmt.Println("Digite um ângulo em radianos:")
	fmt.Scanln(&angulo)
	return angulo
}

func obterN() int {
	var N int
	fmt.Println("Digite um valor inteiro para N:")
	fmt.Scanln(&N)
	return N
}

func calcularFatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calcularFatorial(n-1)
}
