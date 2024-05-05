package main

import "fmt"

func main() {
    var n, m int
    fmt.Scanln(&n)

    v := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanln(&v[i])
    }

    fmt.Scanln(&m)
    teste := make([]int, m)
    for i := 0; i < m; i++ {
        fmt.Scanln(&teste[i])
    }
    
     
    for i:=0;i<m;i++{
        found:=false
        for j:=0;j<n;j++{
            
            if v[j] == teste[i]{
                
                fmt.Println("ACHEI")
                found = true
                break
                
                
            } 
        }
        if !found{
            fmt.Println("NÃO ACHEI")
        }
    }

}
