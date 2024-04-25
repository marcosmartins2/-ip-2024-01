package main

import "fmt"

func main() {

	var linhas, colunas int

	fmt.Println("Digite o número de linhas:")
	fmt.Scanln(&linhas)
	fmt.Println("digite o numero de colunas:")
	fmt.Scanln(&colunas)

	for i := 2; i <= linhas; i++ {

		fmt.Printf("(%v,1)", i)

		for j := 2; j <= colunas; j++ {

			if j < i {

				fmt.Printf("-(%v,%v)", i, j)
			}

		}
		fmt.Println()
	}
}
