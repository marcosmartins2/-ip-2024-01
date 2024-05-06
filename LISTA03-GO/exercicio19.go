package main

import "fmt"

var n int

func main() {

	fmt.Scanln(&n)

	vetor := make([]int, n)

	for i := 0; i < len(vetor); i++ {

		fmt.Scanln(&vetor[i])

	}
	ordenarnumeros(vetor)
	fmt.Println("-------------------")

	numero := make(map[int]int)

	for _, valor := range vetor {

		numero[valor]++

		if numero[valor] == 1 {

			fmt.Println(valor)
		}

	}

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
