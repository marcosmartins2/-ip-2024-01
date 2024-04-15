package main

import (
	"fmt"
	"math"
)

var h, a, areadabase, v float64

func main() {

	fmt.Println("Digite a altura da pirâmide(em metros):")
	fmt.Scanln(&h)
	fmt.Println("Digite o comprimento da aresta do hexagono que forma a pirâmide(em metros):")
	fmt.Scanln(&a)

	areadabase = (a * a * 3 * math.Sqrt(3)) / 2
	v = (areadabase * h) / 2

	fmt.Println("O volume da pirâmide é:", v, "metros cubicos")

}
