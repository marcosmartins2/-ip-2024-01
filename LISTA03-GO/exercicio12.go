package main

import (
	"fmt"
)

var n int

func main() {

	fmt.Scanln(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Scanln(&vetor[i])

	}
	ordenar := ordenarnumeros(vetor)
	fmt.Println(ordenar)
}

func ordenarnumeros(vetor []int) []int {

	for i := 1; i < n; i++ {

		for j := 0; j < n; j++ {

			if vetor[j] > vetor[i] {

				vetor[j], vetor[i] = vetor[i], vetor[j]

			}

		}

	}

	return vetor

}
