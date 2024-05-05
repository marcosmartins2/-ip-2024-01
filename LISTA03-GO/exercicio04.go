package main

import "fmt"

var n int

func main() {

	fmt.Scanln(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Scanln(&vetor[i])

	}
	fmt.Println(vetor)

}
