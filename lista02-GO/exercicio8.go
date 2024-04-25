package main

import "fmt"

func main() {
	var times int
	fmt.Println("Digite a quantidade de times que disputarão o campeonato:")
	fmt.Scanln(&times)

	if times < 2 {
		fmt.Println("Campeonato inválido!")
		return
	}

	fmt.Println("Finais possíveis:")

	for i := 1; i <= times; i++ {
		for j := i + 1; j <= times; j++ {
			fmt.Printf("Final: Time%d X Time%d\n", i, j)
		}
	}
}
