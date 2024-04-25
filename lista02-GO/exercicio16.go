package main

import (
	"fmt"
	"math"
)

var n, hipotenusa, cateto1, cateto2 int

func main() {

	fmt.Println("Digite um número inteiro:")
	fmt.Scanln(&n)

	for hipotenusa = 1; hipotenusa <= n; hipotenusa++ {

		for cateto1 = 1; cateto1 < hipotenusa; cateto1++ {

			cateto2 := math.Sqrt(float64(hipotenusa*hipotenusa - cateto1*cateto1))

			if cateto2 == float64(int(cateto2)) && int(cateto2) > cateto1 {

				fmt.Printf("hipotenusa = %v, catetos %v e %v \n", hipotenusa, cateto1, cateto2)
			}

		}

	}

}
