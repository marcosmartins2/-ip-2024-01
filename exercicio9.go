package main

import "fmt"

var A, B, C, delta float64

func main() {

	fmt.Println("Digite o valor do coeficiente A:")
	fmt.Scanln(&A)
	fmt.Println("Digite o valor do coeficiente B:")
	fmt.Scanln(&B)
	fmt.Println("Digite o valor do coeficiente C:")
	fmt.Scanln(&C)

	delta = B*B - 4*A*C

	fmt.Printf("o valor de delta é:%.2f", delta)

}
