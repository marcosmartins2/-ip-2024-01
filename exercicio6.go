package main

import "fmt"

var t1, n, nfor, f float64

func main() {

	fmt.Println("digite o numero de temperaturas a serem convertidas para Celsius:")
	fmt.Scanln(&n)

	nfor = 0

	for nfor < n {

		nfor = nfor + 1

		fmt.Println("digite a temperatura a ser convertida:")
		fmt.Scanln(&t1)

		f = (5 * (t1 - 32)) / 9

		fmt.Println(t1, "FAHRENHEIT EQUIVALE A ", f, "CELSIUS")
	}
}
