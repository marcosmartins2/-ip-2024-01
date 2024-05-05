package main

import "fmt"

var n, k, cont, cont2 int

func main() {

	fmt.Scan(&n, &k)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])

		if vetor[i] <= 0 {
			cont++
		} else {
			cont2++
		}
	}

	if cont >= k {
		fmt.Println("NAO")
		for k := n - 1; k < n && k >= 0; k-- {

			if vetor[k] <= 0 {
				fmt.Println(k + 1)
			}
		}

	} else {
		fmt.Println("SIM")
	}

}
