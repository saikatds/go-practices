package main

import "fmt"

/**
	fibonacci series example : 0 1 1 2 3 5 8 13 21 .....
 */

func fibonacci () func() int {
	i:= 0
	j:= 1

	/**
		this is an anonymous function
	 */
	return func() int {
		k:= i+j
		i = j
		j = k
		return k
	}
}

func main(){
	n:= 10
	nextFibonacciNmbr := fibonacci()
	fmt.Println("Printing Fibonacci Numbers : ")
	fmt.Print(0 ," ") //printing the first two fibonacci numbers
	fmt.Print(1 , " ")
	for i:= 0; i<n ; i++  {
		fmt.Print(nextFibonacciNmbr(), "  ")
	}

}
