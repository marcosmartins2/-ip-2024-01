package main

import "fmt"

var n, njogos, np, pcp, pcg, pca, pcc float64

func main() {

	fmt.Println("digite o número de jogos a serem jogados:")
	fmt.Scanln(&n)

	njogos = 1

	for njogos <= n {
		njogos = njogos + 1

		fmt.Println("digite o numero total de pessoas:")
		fmt.Scanln(&np)
		fmt.Println("digite a porcentagem de pessoas na categoria popular:")
		fmt.Scanln(&pcp)
		fmt.Println("digite a porcentagem de pessoas na categoria geral")
		fmt.Scanln(&pcg)
		fmt.Println("digite a porcentagem de pessoas na categoria arquibancada:")
		fmt.Scanln(&pca)
		fmt.Println("Digite a porcentagem de pessoas na categoria cadeiras:")
		fmt.Scanln(&pcc)

		popular := np * (pcp / 100)

		geral := np * (pcg / 100)

		arquibancada := np * (pca / 100)

		cadeiras := np * (pcc / 100)

		vpp := popular * 1

		vpa := arquibancada * 10

		vpg := geral * 5

		vpc := cadeiras * 20

		vtt := vpp + vpa + vpg + vpc

		if pcp+pcg+pca+pcc == 100 {
			fmt.Println("o valor total pago é:", vtt, "reais")

		} else {
			fmt.Println("numeros inválidos!!")
		}
	}
}
