package main

import "fmt"

var n float64

func main() {

	fmt.Println("Digite sua nota:")
	fmt.Scanln(&n)

	if n >= 0.0 && n < 6.0 {

		fmt.Println("NOTA = ", n, "CONCEITO = D")
	}
	if n >= 6 && n < 7.5 {

		fmt.Println("NOTA = ", n, "CONCEITO = C")
	}
	if n >= 7.5 && n < 9 {

		fmt.Println("NOTA = ", n, "CONCEITO = B ")
	}
	if n >= 9 && n <= 10 {

		fmt.Println("NOTA = ", n, "CONCEITO = A")
	}
	if n < 0 || n > 10 {

		fmt.Println("FORMATO INVÁLIDO!!!!")
	}
}
