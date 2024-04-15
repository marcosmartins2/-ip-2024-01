package main

import "fmt"

var x int

func main() {

	fmt.Println("digite um numero inteiro:")
	fmt.Scanln(&x)

	if x%3 == 0 && x%5 == 0 {

		fmt.Println("O número é divisível")
	} else {

		fmt.Println("O número não é divisível")
	}

}
