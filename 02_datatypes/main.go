package main

import "fmt"

func main() {
	//MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	//complex64 complex128
	var age float32= 36.76
	var name = "krunal"
	

	//short hand
	flag:=true
	id,email:=2321,"example@example.com"

	fmt.Println(name, age,flag,id,email)
	fmt.Printf("%T\n", flag)

}
