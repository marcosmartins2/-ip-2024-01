package main

import "fmt"

var n, vetor2 int

func main() {

	for {

		fmt.Scanln(&n)

		if n == 0 {
			return
		}

		vetor := make([]int, n)

		for i := 0; i < n; i++ {

			fmt.Scan(&vetor[i])
		}

		for i := 1; i < n; i++ {

			if vetor[i] > vetor[i-1] {

				vetor2 = vetor[i]
			}

		}
		fmt.Println(vetor2)

	}

}
