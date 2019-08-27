package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type circle struct {
	radius float64
}

type rectangle struct{
	height , width float64
}

func (c circle) area() float64{
	return math.Pi*c.radius*c.radius
}

func (r rectangle) area() float64{
	return  r.height*r.width
}

func getArea(s shape) float64{
	return  s.area()
}

func  main ()  {
	c := circle{3}
	r := rectangle{3 , 4 }

	fmt.Println(" Area of Rectangle : ",getArea(&c))
	fmt.Println("Area of Circle : " , getArea(&r))



}