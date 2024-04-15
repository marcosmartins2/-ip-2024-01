package main

import "fmt"

var h, x1, x2 int
var x3 float64

func main() {

	fmt.Println("Digite o numero de horas que a charrete foi usada:")
	fmt.Scanln(&h)

	if h%3 == 0 {

		x1 = (h / 3) * 10
		fmt.Println("o valor a pagar é:%.2f", x1, "reais")

	} else {

		x2 = h % 3
		x3 = (h/3)*10 + x2*5
		fmt.Println("O VALOR A SER PAGO É:", x3, "REAIS")

	}

}
