package main

import (
	"fmt"
	"math"
	"time"
)

func foo() {
	fmt.Println("I am foo. I am outside of main")
}

func main() {
	foo()
	fmt.Println("I am goLang.")
	fmt.Println("The Time is ", time.Now())
	fmt.Println("Square root of 9 : ", math.Sqrt(9))
}
