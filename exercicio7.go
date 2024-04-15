package main

import "fmt"

var f, c, t, p float64

func main() {

	fmt.Println("digite o valor em Fahrenheit a ser convertido para Celsius:")
	fmt.Scanln(&f)
	fmt.Println("digite o valor em polegada para ser convertido a milimetros:")
	fmt.Scanln(&c)

	t = (5*f - 160) / 9

	p = c * 25.4

	fmt.Println("o seu valor em Celsius é:", t)
	fmt.Println("o seu valor em milimetros é", p)

}
