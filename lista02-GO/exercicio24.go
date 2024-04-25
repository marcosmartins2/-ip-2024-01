package main

import "fmt"

func main() {
	var cosseno float64
	sinal := 1.0
	expoente := 1.0
	angulo := obterAngulo()
	N := obterN()
	for n := 0; n <= N; n++ {
		cosseno += sinal * expoente / float64(calcularFatorial(2*n))
		sinal *= -1
		expoente *= angulo * angulo
	}
	fmt.Printf("cos(%.2f) = %.6f\n", angulo, cosseno)
}

func obterAngulo() float64 {
	var angulo float64
	fmt.Println("Escolha um valor positivo para o ângulo (em radianos):")
	fmt.Scanln(&angulo)
	if angulo >= 0 {
		return angulo
	}
	fmt.Println("O ângulo deve ser positivo. Tente novamente.")
	return obterAngulo()
}

func obterN() int {
	var N int
	fmt.Println("Escolha um valor entre 1 e 9 para N:")
	fmt.Scanln(&N)
	if N >= 1 && N <= 9 {
		return N
	}
	fmt.Println("O valor deve estar entre 1 e 9. Tente novamente.")
	return obterN()
}

func calcularFatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * calcularFatorial(n-1)
}
