package main

import (
	"fmt"
	"math"
	"reflect"
)

/**
Function taking an argument as a function
 */
func compute(fn func(x , y float64 ) float64) float64{
	return fn(3,4)
}

func main()  {
	/**
	declaring function as a variable
	 */
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))

	fmt.Println("Type of Compute return " ,reflect.TypeOf(compute(hypot) ))

	fmt.Println(compute(math.Pow)) // here math.Pow is also a function and we are using it as a argument to the func compute
}
