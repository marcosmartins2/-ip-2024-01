package main

import "fmt"

var n1, n2, n3 float64

func main() {

	fmt.Println("digite sua primeira nota (n1):")
	fmt.Scanln(&n1)
	fmt.Println("digite sua primeira nota (n2):")
	fmt.Scanln(&n2)
	fmt.Println("digite sua primeira nota (n3):")
	fmt.Scanln(&n3)

	media := (n1 + n2 + n3) / 3
	fmt.Println("A média é:", media)

	if media >= 6.0 {
		fmt.Println("aprovado")

	} else {
		fmt.Println("reprovado")
	}

}
