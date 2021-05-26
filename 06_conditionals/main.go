package main

import "fmt"

func main() {
	x := 15
	y := 10
 
	//if else
	if x <= y {
		fmt.Printf("%d is less than %d\n",x,y)

	}else { 
		fmt.Printf("%d is less than %d\n",y,x)	
	}


	color:="grey"
	if color=="red"{
		fmt.Println("color is red")
	} else if color=="blue"{
		fmt.Println("color is blue")
	}else{
		fmt.Println("neither of blue or red")
	}


	switch color{
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("neither of blue or red")



		
	}
}

//else block should be present immediately after if block ends
