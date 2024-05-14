package main

import "fmt"

var q1, q2, n int

func main() {

	fmt.Println("Digite o tamanho da sequencia 1:")
	fmt.Scanln(&q1)
	fmt.Println("Digite o tamanho da sequencia 2:")
	fmt.Scanln(&q2)

	slice1 := make([]int, q1)
	slice2 := make([]int, q2)

	for i := 0; i < q1; i++ {
		fmt.Scanln(&slice1[i])
	}
	for i := 0; i < q2; i++ {
		fmt.Scanln(&slice2[i])
	}
	ordenarnumeros(slice1)
	ordenarnumeros(slice2)
	fmt.Println("----------------")
	j := 0
	k := 0
	continuar := true
	slice3 := make([]int, 0)
	fmt.Println(slice3)
	for continuar {

		if j < len(slice1) && (k >= len(slice2) || slice1[j] < slice2[k]) {
			slice3 = append(slice3, slice1[j])
			j++
			if j == len(slice1) && k == len(slice2) {
				continuar = false
			}
			continue
		} else if k < len(slice2) {
			slice3 = append(slice3, slice2[k])
			k++
			if k == len(slice2) && j == len(slice1) {
				continuar = false
			}
			continue
		}

	}

	fmt.Println(slice3)

}

func ordenarnumeros(vetor []int) []int {
	n = len(vetor)
	for i := 1; i < n; i++ {

		for j := 0; j < n; j++ {

			if vetor[j] > vetor[i] {

				vetor[j], vetor[i] = vetor[i], vetor[j]

			}

		}

	}

	return vetor

}

