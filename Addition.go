package main

import "fmt"

func add(a float64, b float64) float64 {
	return a + b
}

func strManipulation(a, b string) string {
	return a + b
}

func main() {
	var a = 5.6
	var b = 9.1

	fmt.Println("Addition result : ", add(a, b))

	fmt.Println("String functions : ", strManipulation("Saikat", " Dey Sarkar"))
}
