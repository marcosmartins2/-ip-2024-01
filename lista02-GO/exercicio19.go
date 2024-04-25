package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Println("Digite um numero inteiro maior que 0:")
	fmt.Scan(&m)

	for n := 1; n <= m; n++ {
		var pivo, inicio, fim, sum int
		if n%2 == 1 { // Se n é ímpar
			pivo = (n * n * n) / n
			inicio = pivo - (n/2)*2
			fim = pivo + (n/2)*2
		} else { // Se n é par
			pivo = (n * n * n) / n
			inicio = pivo - (n/2)*2 + 1
			fim = pivo + (n/2)*2 - 1
		}

		fmt.Printf("%d * %d * %d = ", n, n, n)
		for i := inicio; i <= fim; i += 2 {
			fmt.Printf("%d", i)
			sum += i
			if i != fim {
				fmt.Printf(" + ")
			}
		}
		fmt.Printf(" = %d\n", sum)
	}
}
