package main

import "fmt"

var matricula int
var notafinal, n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11, n12, n13, nt, p, mediaprovas float64

func main() {

	fmt.Print("Digite o número da sua matrícula,as suas 8 notas da prova,as suas 5 notas nas listas de exercicios,a nota do trabalho final e a quantidade de horas assistidas:")
	fmt.Scanln(&matricula, &n1, &n2, &n3, &n4, &n5, &n6, &n7, &n8, &n9, &n10, &n11, &n12, &n13, &nt, &p)

	mediaprovas = (n1 + n2 + n3 + n4 + n5 + n6 + n7 + n8) / 8
	notafinal = 0.7*mediaprovas + 0.15*(n9+n10+n11+n12+n13)/5 + 0.15*nt

	if matricula == -1 && n1 == -1 && n2 == -1 && n3 == -1 && n4 == -1 && n5 == -1 && n6 == -1 && n7 == -1 && n8 == -1 && n9 == -1 && n10 == -1 && n11 == -1 && n12 == -1 && n13 == -1 && nt == -1 && p == -1 {
		fmt.Println("Fim da entrada.")
		return
	}

	switch {

	case notafinal >= 6 && p >= 96:
		fmt.Println("matrícula:", "", matricula, ",", "Nota final:", "", notafinal, ",", "situação final: aprovado")
	case notafinal >= 6 && p < 96:
		fmt.Println("matrícula:", "", matricula, ",", "Nota final:", "", notafinal, ",", "situação final: reprovado por presença")
	case notafinal < 6 && p > 96:
		fmt.Println("matrícula:", "", matricula, ",", "Nota final:", "", notafinal, ",", "situação final: reprovado por nota")
	case notafinal < 6 && p < 96:
		fmt.Println("matrícula:", "", matricula, ",", "Nota final:", "", notafinal, ",", "situação final: reprovado por nota e por presença")
	default:
		fmt.Println("formato inválido!!!")

	}

}
