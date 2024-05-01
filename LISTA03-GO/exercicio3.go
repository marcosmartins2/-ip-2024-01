package main


import"fmt" 


var n int
func main() {
    
    fmt.Scanln(&n)
    
    vetor := make([]int,n)
    
    for i:=0;i<n;i++{
        
        fmt.Scan(&vetor[i])
        
    }
    
    for i:=n-1; i>=0 ; i--{
        
        
        fmt.Printf("%v ",vetor[i])
    }
    
    
    
    
    
}
