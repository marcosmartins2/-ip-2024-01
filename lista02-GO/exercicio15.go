package main

import "fmt"

var T, i, c1, d1, u1, c2, d2, u2, n1, n2, inversao1, inversao2 int

func main() {

	fmt.Println("digite o número de vezes que deseja rodar o programa:")
	fmt.Scanln(&T)

	i := 0
	for i = 0; i < T; i++ {

		fmt.Println("Digite o primeiro número:")
		fmt.Scanln(&n1)
		fmt.Println("Digite o segundo número:")
		fmt.Scanln(&n2)

		c1 = n1 / 100
		d1 = (n1 % 100) / 10
		u1 = (n1 % 100) % 10

		c2 = n2 / 100
		d2 = (n2 % 100) / 10
		u2 = (n2 % 100) % 10

		inversao1 = c1*1 + d1*10 + u1*100

		inversao2 = c2*1 + d2*10 + u2*100

		if inversao1 > inversao2 {

			fmt.Println(inversao1)
		} else {

			fmt.Println(inversao2)
		}

	}

}
