package main

import "fmt"

var n, x, z int

func main() {

	fmt.Println("Digite um numero inteiro N (maior que 5 e menor que 2000):")
	fmt.Scanln(&n)

	if n < 5 || n > 2000 {

		fmt.Println("Valor inválido:")
	}

	x = 0

	for x <= n {

		if (x+2)%2 == 0 {

			z = x * x
			fmt.Println(x, z)

		}
		x++
	}

}
