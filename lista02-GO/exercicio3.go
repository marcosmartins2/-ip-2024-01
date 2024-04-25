package main

import "fmt"

var n, z, y int

func main() {

	fmt.Println("Digite um número n:")
	fmt.Scanln(&n)
	y = 1

	for z = 1; z <= n; z++ {
		y *= z

	}
	fmt.Println(y)

}
