package main

import "fmt"

var n, x, z, y float64

func main() {

	fmt.Println("Digite um numero(inteiro):")
	fmt.Scanln(&n)

	if n < 1 {
		fmt.Println("formato invalido!!!!!")
	}
	y = 0

	for x = 1; x <= n; x++ {

		y += 1 / x

	}
	fmt.Println(y)

}
