package main

import "fmt"

var n, soma int

func main() {

	fmt.Scanln(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Scanln(&vetor[i])

	}

	for i := 1; i <= n; i++ {

		soma += vetor[i-1]

	}
	fmt.Println(soma)
}
