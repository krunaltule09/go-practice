package main

import "fmt"

func main() {
    ids:=[]int{34,67,23,64,453,34}
    

	for i,ID:=range ids{
        fmt.Printf("%d-ID:%d\n",i,ID)
    }
    sum:=0
    for _,id:=range ids{
        sum+=id
    }

    fmt.Println("sum = ",sum)

}
