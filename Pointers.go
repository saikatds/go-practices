package main

import "fmt"

func main(){
	x := "A" // This is dynamic
	a := &x // To become a pointer of a variable we need to add '&' with the variable

	fmt.Println("a : ",a) // here it will print the memory location of a as a stores the memory location.
	fmt.Println("*a : ",*a) // Prints the value

	*a = "B"

	fmt.Println("x : ",x) // accessing the memory location of X using * and then changing the value of it

	*a = *a + *a // String concatination the updated value with the same

	fmt.Println("x : ",x) // print the value

	fmt.Println("&x : ",&x)// print the memory location

	fmt.Println("*a : ",*a)//print the value

	fmt.Println("a : ",a)//print the memory location


}
