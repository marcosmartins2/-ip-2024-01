package main

import "fmt"

var n int

func main() {

	fmt.Println("Digite um número para fatorá-lo em fatores primos:")
	fmt.Scanln(&n)

	valores := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

	j := 0

	var i int

	m := n

	armazenar := 0

	if n <= 0 {
		fmt.Println("Não é possivel para o numero ", n)
		return
	}

	fmt.Printf("%v = ", m)

	for i = 0; i <= n; i++ {

		if n%valores[j] == 0 {

			n = n / valores[j]

			armazenar = valores[j]

			fmt.Printf(" %v ", armazenar)
			break

		}
		j++
	}

	for i = i; i <= n; i++ {
		if n%valores[j] == 0 {

			n = n / valores[j]

			armazenar = valores[j]

			fmt.Printf(" x %v ", armazenar)

		} else {
			j++
		}

	}

}
