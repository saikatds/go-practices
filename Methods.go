package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x , y float64
}

/**
We can also define types other than struct . like here we have defined float64 as a special type
 */
type MyFloat float64

/**
	As this is the struct so changing any value of the struct will not affect out of its scope
	It will not the values actually.
 */
func (v vertex) abs () float64 {
	return math.Sqrt(v.x*v.x+v.y*v.y)
}

/**
	As this is the pointer to the struct it literally change the values of the struct and
	it will remain as the changed values.
 */
func (v *vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (f MyFloat) getAbs() float64{
	if f < 0 {
		return float64(-f)
	}
	return  float64(f)
}

func  main()  {
	v:= vertex{3, 4}
	fmt.Println(v.abs())
	v.Scale(10)
	fmt.Println(v.abs())

	f:= MyFloat(-2)
	fmt.Println(f.getAbs())

}