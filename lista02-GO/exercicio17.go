package main

import "fmt"

var codigo, qtd int
var compra, venda float64

func main() {

	var lucro10, lucro1020, lucro20, maiorlucro, maiormercadoria float64 = 0, -1, 0, 0, 0
	var vtc, vtv, pl float64
	var maiorcodigo, maiorcodigo2 int = 0, 0

	for {

		fmt.Println("Digite o código da mercadoria,seu preço de compra,seu preço de venda e o número de unidades que foram vendidas(para parar de dar entradas digite 0 em algum campo):")
		fmt.Scanln(&codigo, &compra, &venda, &qtd)

		var lucro float64 = venda*float64(qtd) - compra*float64(qtd)

		if venda < 1.1*compra {

			lucro10++

		}
		if venda >= 1.1*compra && venda <= 1.2*compra {

			lucro1020++

		}

		if venda > 1.2*compra {

			lucro20++

		}
		if codigo == 0 || compra == 0 || venda == 0 || qtd == 0 {

			break
		}
		if lucro > maiorlucro {

			maiorlucro = lucro
			maiorcodigo = codigo
		}
		if float64(qtd) > maiormercadoria {

			maiormercadoria = float64(qtd)
			maiorcodigo2 = codigo

		}

		vtc += compra * float64(qtd)
		vtv += venda * float64(qtd)
		pl = ((vtv - vtc) / vtc) * 100

	}

	fmt.Println("Quantidade de mercadorias que geraram lucro menor que 10%:", lucro10)
	fmt.Println(`Quantidade de mercadorias que geraram lucro maior ou igual a 10 % e  menor ou igual a 20%:`, lucro1020)
	fmt.Println("Quantidade de mercadorias que geraram lucro maior que 20%:", lucro20)
	fmt.Println("Código da mercadoria que gerou maior lucro:", maiorcodigo)
	fmt.Println("Código da mercadoria mais vendida:", maiorcodigo2)
	fmt.Printf("Valor total de compras: %.2f , valor total de vendas: %.2f e percentual de lucro total: %.2f", vtc, vtv, pl)

}
