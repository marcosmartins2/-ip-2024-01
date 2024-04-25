package main

import "fmt"

var n1, n2, n3, divisor int

func main() {

	fmt.Println("Digite 3 números inteiros:")
	fmt.Scanln(&n1, &n2, &n3)
	maior := maiornumero(n1, n2, n3)

	mmc := 1

	for i := 2; i <= maior; i++ {

		for {
			if n1%i == 0 || n2%i == 0 || n3%i == 0 {

				divisor = i

			} else {
				break
			}
			if divisor == i {
				fmt.Printf("%v %v %v :%v\n ", n1, n2, n3, divisor)
				mmc *= divisor
			}
			if n1%i == 0 {

				n1 = int(n1 / i)
			}
			if n2%i == 0 {

				n2 = int(n2 / i)
			}
			if n3%i == 0 {

				n3 = int(n3 / i)
			}

		}

	}
	fmt.Printf("MMC: %v", mmc)
}

func maiornumero(n1, n2, n3 int) int {

	maiornumero := 0

	switch {

	case n1 >= n2 && n1 >= n3:
		maiornumero = n1
	case n2 >= n1 && n2 >= n3:
		maiornumero = n2
	case n3 >= n2 && n3 >= n1:
		maiornumero = n3

	}
	return maiornumero

}
