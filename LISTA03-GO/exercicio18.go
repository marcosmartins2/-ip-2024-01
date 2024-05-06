package main

import "fmt"

var n, somab1, somab2 int

func main() {

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		somab1 = 0
		somab2 = 0
		cont2 := 9
		cont := 0

		vetor := make([]int, 11)

		for j := 0; j < 11; j++ {

			fmt.Scan(&vetor[j])
		}
		// vetor na posição 9 é o b1
		//vetor na posição 10 é o b2

		for k := 0; k < 9; k++ {
			cont++
			somab1 += vetor[k] * cont

		}
		for l := 0; l < 9; l++ {

			somab2 += vetor[l] * cont2
			cont2--

		}
		if vetor[9] == somab1%11 && vetor[10] == somab2%11 {

			fmt.Println("CPF válido")

		} else if vetor[9] == 0 && somab1%11 == 10 && vetor[10] == somab2%11 {
			fmt.Println("CPF válido")

		} else if vetor[9] == 0 && somab1%11 == 10 && vetor[10] == 0 && somab2%11 == 10 {
			fmt.Println("CPF válido")

		} else if vetor[9] == somab1%11 && vetor[10] == 0 && somab2%11 == 10 {
			fmt.Println("CPF válido")

		} else {
			fmt.Println("CPF inválido")
		}

	}

}
