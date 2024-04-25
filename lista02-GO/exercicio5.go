package main

import "fmt"

var n, e, controle, slice, contador, contador2 int

func main() {

	fmt.Println("Digite a quantidade de números que o seu programa possuirá:")
	fmt.Scanln(&n)

	numeros := make([]int, n)

	fmt.Println("Digite os elementos do seu programa:")

	for slice := 0; slice < n; slice++ {

		fmt.Scan(&numeros[slice])

	}
	contador = 0
	contador2 = 0
	for controle = 1; controle < n; controle++ {
		if numeros[controle-1] <= numeros[controle] {
			contador++
		} else {
			if contador > contador2 {
				contador2 = contador
			}
			contador = 0
		}

	}

	if contador > contador2 {

		contador2 = contador
	}
	fmt.Println("o maior valor de segmento é:", contador2)
}
