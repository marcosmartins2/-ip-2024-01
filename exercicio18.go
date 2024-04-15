package main

import "fmt"

var n1, r, n, sn float64

func main() {

	fmt.Println("Digite o numero inicial(n1)da sua progressão aritmetica:")
	fmt.Scanln(&n1)
	fmt.Println("Digite a razão da sua progressão aritmética:")
	fmt.Scanln(&r)
	fmt.Println("Digite o numero de termos da sua progressão aritmetica:")
	fmt.Scanln(&n)

	sn = (2*n1 + r*n - r) * n / 2

	fmt.Println("A soma da sua progressão aritmética é:", sn)

}
