package main

import "fmt"

func main() {
	num1:=10
	ptr:=&num1
	fmt.Println(ptr)
	fmt.Printf("%T\n",ptr)
	fmt.Println(*ptr,*&num1)
	*ptr=90
	fmt.Println(num1)

	

}
