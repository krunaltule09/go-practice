package main

import "fmt"

func greet(name string) string{
	return "Hello "+name
}

func getSum(num1 , num2 int) int{
	return num1+num2
}

func main(){
	fmt.Println(greet("krunal"))
	fmt.Println(getSum(9,7))
}