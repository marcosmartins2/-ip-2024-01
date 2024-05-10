package main 


import"fmt"
    
var n,cont int



func main(){
    
    fmt.Scanln(&n)
    
    slice := make([]int,n)
    
    
    for i:=0;i<n;i++{
        
        
        fmt.Scan(&slice[i])
    }
    
    frequencia := make(map[int]int)
    for _, num := range slice {
        frequencia[num]++
    }
    for i:=0;i<n;i++{
        if frequencia[i] == 1{
            cont++
        }
    }
    fmt.Println(cont+1)
    }
    
