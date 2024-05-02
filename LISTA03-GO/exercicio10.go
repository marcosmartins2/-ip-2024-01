package main

import "fmt"

var n int

func main() {
	cont := 0
	vetor2 := 0
	i2 := 0

	fmt.Scanln(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Scanln(&vetor[i])

	}
	for j := 0; j < n; j++ {

		if vetor[j] == vetor[n-1] {

			cont++
		}

	}
	vetor2 = 0
	for k := 0; k < n; k++ {

		if vetor2 < vetor[k] {

			vetor2 = vetor[k]
			i2 = k

		}

	}

	fmt.Printf("NOTA %v, %v vezes\n NOTA %v, índice %v", vetor[n-1], cont, vetor2, i2)

}

