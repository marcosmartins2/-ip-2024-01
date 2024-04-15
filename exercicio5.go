package main

import "fmt"

var n int
var c float64
var tipo string
var comercial float64
var industrial float64
var residencial float64

func main() {

	fmt.Println("Digite o numero da sua conta:")
	fmt.Scanln(&n)

	fmt.Println("Digite a quantidade de agua consumida em metros cubicos:")
	fmt.Scanln(&c)

	fmt.Println("Digite o seu tipo de consumidor C(comercial),I(industrial),R(residencial):")
	fmt.Scanln(&tipo)

	fmt.Println("CONTA = ", n)

	if tipo == "C" {
		if c > 80 {
			comercial = (c-80)*0.25 + 500
			fmt.Println("VALOR DA CONTA = ", comercial)
		} else {
			fmt.Println("VALOR DA CONTA = 500 reais")
		}
	}
	if tipo == "I" {
		if c > 100 {
			industrial = (c-100)*0.04 + 800
			fmt.Println("VALOR DA CONTA =", industrial, "reais")
		} else {
			fmt.Println("VALOR DA CONTA = 800 reais")
		}
	}
	if tipo == "R" {
		residencial = (c * 0.05) + 5
		fmt.Println("VALOR DA CONTA =", residencial, "reais")
	}

	if tipo != "I" && tipo != "C" && tipo != "R" {

		fmt.Println("formato inválido!!!")
	}
}
