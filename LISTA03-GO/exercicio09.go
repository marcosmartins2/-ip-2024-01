package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var soma float64 
	fmt.Scanln(&n)

	matriz := make([][]float64, n) // Cria um slice de slices com tamanho n

	for i := 0; i < n; i++ {
		matriz[i] = make([]float64, 3) // Inicializa cada slice interna com tamanho 3
		for j := 0; j < 3; j++ {
			fmt.Scan(&matriz[i][j]) // Lê os valores para cada elemento da matriz

		}
	}

	for l := 0; l < n-1; l++ { // realizar as operações com as matrizes
		soma = 0
		  
		  x:=(matriz[l][0] - matriz[l+1][0])*(matriz[l][0] - matriz[l+1][0])
		  y:=(matriz[l][1] - matriz[l+1][1])*(matriz[l][1] - matriz[l+1][1])
		  z:=(matriz[l][2] - matriz[l+1][2])*(matriz[l][2] - matriz[l+1][2])

			soma = math.Sqrt(x+y+z)
			fmt.Println("-----------------------")
	    	fmt.Printf("%.2f\n", soma)


		}
	
	}


