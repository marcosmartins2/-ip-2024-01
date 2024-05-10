package main

import "fmt"

var n int

func main() {

	fmt.Scanln(&n)

	slice := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Scanln(&slice[i])

	}
	crescente(slice)
	fmt.Println("---------------")

	for j := 0; j < n; j++ {

		if slice[j]%2 == 0 {
			fmt.Println(slice[j])
		}

	}
	decrescente(slice)

	for k := 0; k < n; k++ {

		if slice[k]%2 == 1 {
			fmt.Println(slice[k])
		}
	}

}
func crescente(vetor []int) []int {

	for i := 1; i < n; i++ {

		for j := 0; j < n; j++ {

			if vetor[j] > vetor[i] {

				vetor[j], vetor[i] = vetor[i], vetor[j]

			}

		}

	}

	return vetor

}
func decrescente(vetor []int) []int {

	for i := 1; i < n; i++ {

		for j := 0; j < n; j++ {

			if vetor[j] < vetor[i] {

				vetor[j], vetor[i] = vetor[i], vetor[j]

			}

		}

	}

	return vetor

}
