package main

import "fmt"

var n, vetor2, vetor3 int

func main() {

	fmt.Scanln(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Scan(&vetor[i])

	}

	for j := n - 1; j < n && j >= 0; j-- {
		fmt.Printf("%v ", vetor[j])

	}
	vetor2 = vetor[0]
	for k := 1; k < n; k++ {

		if vetor2 < vetor[k] {

			vetor2 = vetor[k]

		}

	}
	vetor3 = vetor[0]
	for l := 1; l < n; l++ {

		if vetor3 > vetor[l] {

			vetor3 = vetor[l]
		}

	}
	fmt.Printf("\n %v \n %v ", vetor2, vetor3)

}
