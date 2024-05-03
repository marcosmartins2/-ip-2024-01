package main

import (
	"fmt"
	"sort"
)

var n int

func main() {

	fmt.Scanln(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Scanln(&vetor[i])

	}
	sort.Ints(vetor)
	fmt.Println("-----------------")

	for j := 0; j < n; j++ {

		fmt.Println(vetor[j])

	}
}
