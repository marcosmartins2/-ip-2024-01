package main

import (
	"fmt"
	"math"
)

var n, e float64

func main() {

	fmt.Println("digite o número n, que se deseja obter a raiz quadrada")
	fmt.Scanln(&n)
	r := 1.0
	erro := n

	for {
		fmt.Println("digite o erro que deverá ser considerado")
		fmt.Scanln(&e)
		if e < n {
			fmt.Println("o erro admitido deve ser menor que o número n")
			break
		}

	}
	for erro > e {
		r = (r + (n / r)) / 2

		erro = math.Abs(n - (r * r))
		fmt.Printf("r: %.9f, erro: %.9f\n", r, erro)
	}

}
