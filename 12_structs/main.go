package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	fname  string
	lname  string
	city   string
	age    int
	gender string
}

//Greeting method (value receiver)
func (p Person) greet() string {
	return "hello my name is " + p.fname + " " + p.lname + " i am " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday(){
	p.age++
}

func (p1 *Person) getMarried(p2 *Person){
	if(p1.gender=="m" && p2.gender=="f"){
		p2.lname=p1.lname
	}else if(p1.gender=="f" && p2.gender=="m"){
		p1.lname=p2.lname
	}
}

func main() {
	person1 := Person{fname: "krunal", lname: "tule", city: "mumbai", age: 21, gender: "m"}

	person2 := Person{fname: "diana", lname: "daniels", city: "new york", age: 21, gender: "f"}


	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

	person2.getMarried(&person1)

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
