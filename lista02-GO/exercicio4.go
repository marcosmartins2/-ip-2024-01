package main

import "fmt"

var n, K, s, controle float64

var i, icomeço float64

func main() {

	fmt.Println("Digite um numero de 0 a 9:")
	fmt.Scanln(&n)
	fmt.Println("Digite o número que começará a tabuada:")
	fmt.Scanln(&i)
	fmt.Println("Digite o número de fatores requeridos na tabuada:")
	fmt.Scanln(&K)
	fmt.Println("Digite o fator que aumenta a cada iteração da tabuada:")
	fmt.Scanln(&s)

	if n < 0 || n > 9 {

		fmt.Println("FORMATO INVÁLIDO!!!!")

	}

	icomeço = i

	fmt.Println("TABUADA DA SOMA:\n")

	for controle = 1; controle <= K; controle++ {
		fmt.Printf("%v + %v = %.2f\n", n, i, n+i)
		i += s

	}
	i = icomeço
	fmt.Println("TABUADA DA SUBTRAÇÃO:\n")
	for controle = 1; controle <= K; controle++ {

		fmt.Printf("%v - %v = %.2f\n", n, i, n-i)
		i += s

	}
	i = icomeço
	fmt.Println("TABUADA DA MULTIPLICAÇÃO:\n")
	for controle = 1; controle <= K; controle++ {
		fmt.Printf("%v x %v = %.2f\n", n, i, n*i)
		i += s

	}
	i = icomeço
	fmt.Println("TABUADA DA DIVISÃO:\n")
	for controle = 1; controle <= K; controle++ {
		fmt.Printf("%v / %v = %.2f\n", n, i, n/i)
		i += s

	}

}
