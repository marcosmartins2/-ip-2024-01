package main

import "fmt"

var r, h, ac, at, s2, c float64

func main() {

	fmt.Println("digite o valor do raio da lata de cerveja( em metros):")
	fmt.Scanln(&r)
	fmt.Println("digite o valor da altura da lata de cerveja( em metros):")
	fmt.Scanln(&h)
	pi := 3.14159
	ac = 2 * (pi * r * r)
	at = (2 * r * h * pi)
	s2 = ac + at
	c = s2 * 100

	fmt.Println("O custo da lata é:", c, "reais")

}
