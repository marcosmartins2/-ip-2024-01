package main

import "fmt"

var x, y, n, vrp, ckw, vcd float64

func main() {

	fmt.Println("Digite o valor do sálario mínimo:")
	fmt.Scanln(&x)
	fmt.Println("digite a quantidade de kw gasta:")
	fmt.Scanln(&y)
	n = 0.7 * y
	vrp = x * n / 100
	ckw = vrp / y
	fmt.Println("custo por KW:", ckw, "valor em reais pago:", vrp)
	vcd = vrp * 0.9
	fmt.Println("valor com desconto:", vcd)

}
