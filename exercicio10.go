package main

import "fmt"

var a, b, c, d, det float64

func main() {

	fmt.Println("Digite o valor dos coeficientes da primeira linha(separado por vírgulas):")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println("Digite o valor dos coeficientes da segunda linha(separado por vírgulas):")
	fmt.Scanln(&c)
	fmt.Scanln(&d)

	det = a*d - b*c

	fmt.Println("o valor do determinante é:%.2f", det)

}
