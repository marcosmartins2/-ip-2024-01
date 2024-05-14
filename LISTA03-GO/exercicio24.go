package main

import "fmt"

var n int // n é o numero de elementos a serem ordenados

func main() {

	continuar := true

	for continuar {
		fmt.Scan(&n)
		if n == 0 {
			continuar = false
		}

		vetor := make([]int, n)
		for i := range vetor {
			fmt.Scan(&vetor[i])
		}

		vOrd := countingsort(vetor)
		for _, num := range vOrd {
			fmt.Printf("%v ", num)

		}
		fmt.Println()
	}

}

func countingsort(vetor []int) []int {
	maior := 0
	for _, valor := range vetor {

		if valor > maior {

			maior = valor
		}
	}
	vCount := make([]int, maior+1)

	for _, valor := range vetor {

		vCount[valor]++
	}
	for i := 1; i < len(vCount); i++ {
		vCount[i] += vCount[i-1]
	}
	vOrd := make([]int, len(vetor))
	for _, valor := range vetor {
		vOrd[vCount[valor]-1] = valor
		vCount[valor]--
	}

	return vOrd

}
