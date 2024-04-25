package main

import "fmt"

var pa, pb, pa2, pb2, anos float64

func main() {
	fmt.Println("Digite q população do país A:")
	fmt.Scanln(&pa)
	fmt.Println("Digite q população do país B:")
	fmt.Scanln(&pb)

	pa2 = pa
	pb2 = pb

	for anos = 0; pa2 <= pb2; anos++ {
		pa2 *= 1.03
		pb2 *= 1.015

	}
	fmt.Println("anos = ", anos)

}
