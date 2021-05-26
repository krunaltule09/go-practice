package main

import "fmt"

func main() {
	// emails:=make(map[int] string)
	// emails[34]="krunal"
	// emails[12]="bob"

	emails:=map[int] string{23:"krunal",67:"bob",100:"sharon"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails[23])

	//delete element
	delete(emails,67)
	fmt.Println(emails)


}
