package main

import (
	"fmt"
)

// Função para verificar se um elemento está presente em um conjunto
func elementoPresente(conjunto []int, elemento int) bool {
	for _, num := range conjunto {
		if num == elemento {
			return true
		}
	}
	return false
}

func main() {
	var TA, TB int

	// Ler o tamanho do conjunto A
	for {
		fmt.Scanln(&TA)
		if TA >= 1 && TA <= 100 {
			break
		} else {
			fmt.Println("Tamanho inválido para o conjunto A. Digite novamente:")
		}
	}

	// Ler o tamanho do conjunto B
	for {
		fmt.Scanln(&TB)
		if TB >= 1 && TB <= 100 {
			break
		} else {
			fmt.Println("Tamanho inválido para o conjunto B. Digite novamente:")
		}
	}

	// Ler os elementos do conjunto A
	conjuntoA := make([]int, TA)
	for i := 0; i < TA; {
		var elemento int
		fmt.Scan(&elemento)
		if !elementoPresente(conjuntoA, elemento) {
			conjuntoA[i] = elemento
			i++
		} else {
			fmt.Println("Elemento repetido. Digite novamente:")
		}
	}

	// Ler os elementos do conjunto B
	conjuntoB := make([]int, TB)
	for i := 0; i < TB; {
		var elemento int
		fmt.Scan(&elemento)
		if !elementoPresente(conjuntoB, elemento) {
			conjuntoB[i] = elemento
			i++
		} else {
			fmt.Println("Elemento repetido. Digite novamente:")
		}
	}

	// Calcular e imprimir a união dos conjuntos
	fmt.Print("(")
	for i, num := range conjuntoA {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(num)
	}
	for _, num := range conjuntoB {
		if !elementoPresente(conjuntoA, num) {
			fmt.Print(",", num)
		}
	}
	fmt.Println(")")

	// Calcular e imprimir a interseção dos conjuntos
	fmt.Print("(")
	for i, num := range conjuntoA {
		if i > 0 && elementoPresente(conjuntoB, num) {
			fmt.Print(",", num)
		} else if i == 0 && elementoPresente(conjuntoB, num) {
			fmt.Print(num)
		}
	}
	fmt.Println(")")
}
// ler
