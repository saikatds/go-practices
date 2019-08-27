package main

import "fmt"

/**
This is a function closure. It is just a function which is returning an anonymous function
 */
func generateSequence() func() int {
	i:= 0

	return func () int {
		i+= 1
		return i
	}
}

func main(){
	/**
	 // Here we are defining the function to a variable named as seqNumber. As the function
		returns a function so the variable is also a function.
	 */
	seqNumber := generateSequence()

	fmt.Println(seqNumber())
	fmt.Println(seqNumber())
	fmt.Println(seqNumber())

}