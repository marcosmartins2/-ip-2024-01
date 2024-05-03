package main


import (
    "fmt"
    "sort"
    )

var n int
var mediana float64

func main() {
    
    fmt.Println("digite a quantidade de numeros que seu vetor terá:")
    fmt.Scanln(&n)
    
    fmt.Println("digite os elementos do seu vetor:")
    vetor := make([]int,n)
    
    
    for i:=0 ;i<n;i++{
            
        fmt.Scanln(&vetor[i])    
        
    }
    sort.Ints(vetor)
    
    if n%2 == 0{
        
        mediana = float64(vetor[n/2]+vetor[(n/2)-1])/2
    }
    if n%2 == 1 {
        
        
        mediana = float64(vetor[(n/2)])
    }
    fmt.Println("------------------")
    fmt.Printf("%.2f",mediana)
    
    
    
    
}
