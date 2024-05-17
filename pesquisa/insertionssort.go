package main


import"fmt"

func main() {
   
    insertionSort()
   
   
   
}
func insertionSort() {
   
   
    var vetor = []int{2,7,9,1,5,0}
   
    for i:=1;i<len(vetor);i++{
        j:=i
        for j >0 && vetor[j]<vetor[j-1]{
           
           

                 swap(j,vetor)
                 j--
        }
