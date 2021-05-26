package main

import "fmt"

func main(){
	//Arrays
	var arr [2] int
	fruitArr:=[2] string {"apples","oranges"}
	animalArr:=[] string {"tiger","dog","cat","lion"}  


	//Assign values
	arr[0]=56
	arr[1]=34
	fmt.Println(arr)
	fmt.Println(fruitArr)
	fmt.Println(animalArr[1:3])


}