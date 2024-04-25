package main

import "fmt"

var n int

func main() {
	for {
		fmt.Println("Digite a quantidade números desejados:")
		fmt.Scanln(&n)

		if n == 0 {

			break
		}

		array := make([]float64, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&array[i])

		}
		for i := 0; i < n-1; i++ {

			if array[i] > array[i+1] {
				fmt.Println("não é ordenada")
				break

			} else if array[n-2] < array[n-1] {
				fmt.Println("ORDENADA")
				break

			}

		}
	}
}
