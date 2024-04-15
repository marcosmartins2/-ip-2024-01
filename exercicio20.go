package main

import "fmt"

var x, y, b int

func main() {

	fmt.Println("digite um numero x :")
	fmt.Scanln(&x)
	fmt.Println("digite um numero y:")
	fmt.Scanln(&y)

	if x%2 == 0 {

		b = 0
		fmt.Println("")
		for a := 0; a < y; a++ {

			fmt.Println("\n", x+b)
			b += 2

		}

	}

}
