package main


import (
    "fmt"
    "math"
    )
   
    var n,t int
   
func main() {
   
    fmt.Scanln(&n)
   
    var distancia,cont int = 0,0
    var menordistancia int = 0
    for i:=0;i<n;i++{
     
        cont = 0
        fmt.Scanln(&t)
       
        vetor := make([]int,t)
       
       
        for j:=0;j<t;j++{
           
            fmt.Scan(&vetor[j])
           
        }
       
       
         
           
            for p:=1;p<t;p++{
             
             distancia = int(math.Abs(float64(vetor[p]-vetor[p-1])))
              if menordistancia<distancia{
               
                menordistancia = distancia
              }
               
               
               
            }
            cont = (len(vetor)-1)*len(vetor)/2
           
        
       
       
        fmt.Println(distancia,cont)
    }
   
   
   
   
}
