package main

import "fmt"

var N,K int

func main() {
       
        
      
    for {  
        
        if N<1 || N>1000{
          
          fmt.Scanln(&N)
      } else {
          
          break
      }
      
    }
    
    V := make([]int,N)
    
    cont := 0
    for i:=0;i<N;i++{
        
        fmt.Scan(&V[i])
        
        
    }
    
    fmt.Scanln(&K)
    i := 0
    for i = 0 ; i<N;i++{
        
        if V[i]>=K{
            
            cont++
        
    }
}
fmt.Println(cont)


}
