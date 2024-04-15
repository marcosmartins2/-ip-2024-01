package main

import "fmt"

var salario, c1, c2 float64

func main() {

	fmt.Println("Digite o valor do seu salário:")
	fmt.Scanln(&salario)

	if salario <= 300 {

		c1 = salario * 1.5
		fmt.Println("O salário com reajuste é:%.2f", c1, "reais")

	} else {

		c2 = salario * 1.3
		fmt.Println("O salário com reajuste é:%.2f", c2, "reais.")
	}

}
