package main

import "fmt"

func main() {
	fmt.Println("Counting starts")
	for i := 1; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Counting done")
}
