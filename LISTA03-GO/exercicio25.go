package main

import "fmt"

var n, sena, quina, quadra, cont int

func main() {

	fmt.Println("Digite os 6 números sorteados:")

	sorteados := make([]int, 6)

	for j := 0; j < 6; j++ {
		fmt.Scan(&sorteados[j])
	}

	fmt.Println("Digite quantas apostas foram feitas:")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		aposta := make([]int, 6)
		cont = 0

		for k := 0; k < 6; k++ {

			fmt.Scan(&aposta[k])
		}

		for l := 0; l < 6; l++ {

			for m := 0; m < 6; m++ {

				if aposta[m] == sorteados[l] {

					cont++
				}

			}

		}
		if cont == 6 {

			sena++
		} else if cont == 5 {

			quina++
		} else if cont == 4 {
			quadra++
		}

	}
	if sena == 0 {
		fmt.Printf("Não houve acertador para sena\n")
	}
	if sena != 0 {
		fmt.Printf("Houve %v acertador(es) para a sena\n ", sena)
	}
	if quina == 0 {
		fmt.Printf("nao houve acertador para a quina\n")
	}
	if quina != 0 {
		fmt.Printf("houve %v acertador(es) para a quina\n", quina)
	}
	if quadra == 0 {
		fmt.Printf("não houve acertador para a quadra\n ")
	}
	if quadra != 0 {
		fmt.Printf("houve %v acertador(es) para a quadra\n", quadra)
	}
}
