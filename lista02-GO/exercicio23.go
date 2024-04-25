package main

import "fmt"

func main() {
	var escolha, numero, soma1 int

	for {
		fmt.Println("Escolha um número de 1 a 9:")
		fmt.Scanln(&escolha)
		if escolha > 0 && escolha <= 9 {
			break
		}
		fmt.Println("O número deve estar entre 1 a 9. Tente novamente.")
	}

	contador := 0
	for i := 2; contador < escolha; i++ {
		numero = calcularSomaDivisores(i)
		soma1 = calcularSomaDivisores(numero)
		if numero > i && soma1 == i {
			fmt.Printf("(%d,%d)\n", i, numero)
			contador++
		}
	}
}

func calcularSomaDivisores(num int) int {
	var soma int
	for i := 1; i <= int(num/2); i++ {
		if num%i == 0 {
			soma += +i
		}
	}
	return soma
}
