package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var maior, diferença float64 = 0, 0
	fmt.Scanln(&n)

	matriz := make([][]float64, n) // Cria um slice de slices com tamanho n

	for i := 0; i < n; i++ {
		matriz[i] = make([]float64, 3) // Inicializa cada slice interna com tamanho 3
		for j := 0; j < 3; j++ {
			fmt.Scan(&matriz[i][j]) // Lê os valores para cada elemento da matriz

		}
	}

	for l := 0; l < n-1; l++ { // realizar as operações com as matrizes
		maior = 0
		for c := 0; c < 3; c++ {

			diferença = math.Abs((matriz[l][c] - matriz[l+1][c]))

			if diferença > maior {

				maior = diferença
			}

		}
		fmt.Println("-----------------------")
		fmt.Printf("%.2f", maior)

	}

}
