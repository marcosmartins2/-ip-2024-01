package main

import "fmt"

var x, y, z int

func main() {

	fmt.Println("Digite o número que deseja partir:")
	fmt.Scanln(&x)
	fmt.Println("Digite o número de numeros subsequentes a este que vc deseja que apareça:")
	fmt.Scanln(&y)

	if x%2 != 0 {

		fmt.Println("O PRIMEIRO NUMERO NÃO É PAR")
	} else {
		for z := 0; z < y; z++ {

			fmt.Println(x)
			x += 2
		}
	}

}
