package main

import "fmt"

var n1, n2, n3, x, y int

func main() {
	fmt.Println("digite 1 numero:")
	fmt.Scanln(&n1)
	fmt.Println("digite mais um numero:")
	fmt.Scanln(&n2)
	fmt.Println("digite mais um numero:")
	fmt.Scanln(&n3)

	if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
		fmt.Println("formato inválido")
	}

	x = n1*100 + n2*10 + n3
	y = x * x
	fmt.Println(x, ",", y)

}
