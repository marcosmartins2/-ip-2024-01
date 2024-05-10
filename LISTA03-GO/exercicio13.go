package main 


import(
    "fmt"
    "sort"
    )

var n,cont,numero int

func main(){
    var cont2 = 0
    
    fmt.Scanln(&n)
    
    slice := make([]int,n)
    
    
    for i:=0;i<n;i++{
        
        fmt.Scanln(&slice[i])
        
        
        
        
    }
    sort.Ints(slice)
    
    
    for j:=0;j<n;j++{
        cont = 0
        
        for k:=0;k<n;k++{
            
            if slice[j] == slice[k]{
                cont++
            }
            if cont>cont2  {
                cont2 = cont
                numero = slice[j]
            }
        }
        
    }
     fmt.Printf("%v \n%v",numero,cont2)
}
