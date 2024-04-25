package main

import "fmt"

var n, soma int

func main() {

	fmt.Println("Digite um numero inteiro:") // escanear o numero solicitado
	fmt.Scanln(&n)

	cont := 1
	vetor := [999]int{}

	for i := 1; i < n; i++ {

		if n%i == 0 {

			vetor[cont-1] = i
			cont++

		}
	}

	tamanho := len(vetor)

	fmt.Printf("%v = 1", n)

	for i := 0; i < tamanho; i++ {

		if vetor[i] == 0 {
			break
		}

		soma += vetor[i]

	}

	for i := 1; i < tamanho; i++ {

		if vetor[i] == 0 {
			break
		}

		fmt.Printf(" + %v", vetor[i])

	}
	if soma == n {
		fmt.Printf(" = %v (O NUMERO É PERFEITO)", soma)
	} else {

		fmt.Printf(" = %v (O NUMERO NAO É PERFEITO)", soma)
	}

}
