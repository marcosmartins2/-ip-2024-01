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
		ordenada := true

		for i := 0; i < n-1 && ordenada == true; i++ {

			if array[i] > array[i+1] {
				ordenada = false

			}

		}
		if ordenada {
			fmt.Println("ORDENADA")
		} else {
			fmt.Println("DESORDENADA")
		}

	}
}

