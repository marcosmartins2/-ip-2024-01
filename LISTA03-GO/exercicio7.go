package main

import "fmt"

var n, vetor2 int

func main() {

	for {

		fmt.Scanln(&n)

		vetor := make([]int, n)

		for i := 0; i < n; i++ {

			fmt.Scanln(&vetor[i])
		}

		for i := 1; i < n; i++ {

			if vetor[i] > vetor[i-1] {

				vetor2 = vetor[i]

			}

		}
		for cont := 0; cont <= vetor2; cont++ {

			cont2 := 0

			for j := 0; j < n; j++ {

				if vetor[j] <= cont {

					cont2++

				}

			}
			fmt.Printf("(%v) %v\n", cont, cont2)
		}

	}
}
