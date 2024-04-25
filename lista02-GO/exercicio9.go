package main

import "fmt"

var n, controle int

func main() {

	fmt.Println("Digite um número para saber se é primo ou não:")
	fmt.Scanln(&n)

	for controle := 2; controle < n; controle++ {

		if n <= 1 {

			fmt.Println("não é primo")
			break
		}

		if n%controle == 0 {
			fmt.Println("o número não é primo")
			break
		} else {
			fmt.Println("O número é primo")
			break

		}

	}
	if n == 2 {
		fmt.Println("O numero é primo")

	}
}
