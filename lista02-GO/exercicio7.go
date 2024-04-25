package main

import "fmt"

var n, numeros, somaPares, quantPares, quantImpares, somaImpares int
var mediaImpares, mediaPares float64

func main() {

	quantPares = 0
	quantImpares = 0

	for {

		var numeros int
		fmt.Println("Digite os numeros desejados na sequencia(diferentes de zero,utilize zero quando nao quiser mais inserir números na sequencia):")
		fmt.Scanln(&numeros)

		if numeros%2 == 0 {

			quantPares++
			somaPares += numeros
			mediaPares = float64(somaPares) / float64(quantPares)

		} else {

			quantImpares++
			somaImpares += numeros
			mediaImpares = float64(somaImpares) / float64(quantImpares)

		}
		if numeros == 0 {

			break
		}
	}
	fmt.Println("a media dos pares é:", mediaPares)
	fmt.Println("a media dos impares é:", mediaImpares)
}
