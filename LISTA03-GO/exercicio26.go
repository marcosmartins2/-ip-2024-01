package main

import (
	"fmt"
	"sort"
)

var n, soma int

func main() {

	fmt.Println("Digite o numero de casos de testes:")

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {

		anoes := make([]int, 9)

		for l := 0; l < 9; l++ {
			fmt.Scanln(&anoes[l])
		}
		sort.Ints(anoes)

		for j := 0; j < len(anoes); j++ {

			for k := j + 1; k < len(anoes); k++ {
				soma = somar(anoes)
				if (soma - anoes[j] - anoes[k]) == 100 {

					printAnoesExcept(anoes, j, k)

				}

			}
		}

	}

}

func somar(x []int) int {

	soma := 0
	for _, valor := range x {

		soma += valor
	}

	return soma

}

func printAnoesExcept(anoes []int, impostor1, impostor2 int) {
	for i, valor := range anoes {
		if i != impostor1 && i != impostor2 {
			fmt.Println(valor)
		}
	}
}
