package main

import "fmt"

// types of for loop for go language

func main(){

	// type 1

	for i:=1 ; i < 10 ; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
	// type 2

	i := 1

	for i < 10 {
		fmt.Print(i, " " )
		i++
	}

	fmt.Println()

	// type 3

	x := 1

	for{
		fmt.Print(x , " ")
		x+=3
		if x>10 {
			break
		}
	}

	fmt.Println()

	// type 4

	y:=10
	for x:= 1 ; y<20 ; y++{
		fmt.Print(x , " ")
		//y+=5
	}

	fmt.Println()

	// range

	var arr = [] int {4,6,8,10}

	for i , j := range arr{ // range returns two values from the slice. 1. index 2. value of that index
		fmt.Println(i,j)
	}
}


