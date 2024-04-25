package main

import "fmt"

var valor, valorinicial, valorfinal, controle, pessoas, totalPessoas, lucro, melhorLucro, npessoas, ingressos float64

func main() {

	fmt.Println("DIgite o valor do ingressso,o valor inicial e o valor final com que se deseja testar:")
	fmt.Scanln(&valor, &valorinicial, &valorfinal)

	if valorfinal > valorinicial {

		fmt.Println("FORMATO INVÁLIDO")
	}
	pessoas = 120
	melhorLucro = 0
	npessoas = 0
	ingressos = 0

	for controle := valorinicial; controle <= valorfinal; controle++ {

		if controle <= valor {

			totalPessoas = pessoas + (50 * (valor - controle))
			lucro = controle*totalPessoas - (200 + 0.05*totalPessoas)

			fmt.Printf("V: %v, N: %v, L: %v \n", controle, totalPessoas, lucro)

		} else {

			totalPessoas = pessoas + (60 * (valor - controle))
			lucro = controle*totalPessoas - (200 + 0.05*totalPessoas)

			fmt.Printf("V: %v, N: %v, L: %v \n", controle, totalPessoas, lucro)

		}
		if lucro > melhorLucro {

			melhorLucro = lucro
			npessoas = totalPessoas
			ingressos = controle

		}

	}
	fmt.Printf("Melhor valor final: %v \n lucro: %v \n numero de ingressos: %v \n", ingressos, melhorLucro, npessoas)

}
