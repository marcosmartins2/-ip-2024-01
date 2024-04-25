package main

import "fmt"

var horas, valor, salario float64
var matricula int

func main() {

	for {
		fmt.Println("Digite o número de matrícula, a quantidade de horas trabalhadas e o valor recebido por hora:")
		fmt.Scanln(&matricula, &horas, &valor)

		if matricula != 0 && horas != 0 && valor != 0 || valor != 0.0 {
			salario = horas * valor

			fmt.Println(matricula, "", salario)

		} else {
			return
		}

	}
}
