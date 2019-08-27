package main

import "fmt"


func main(){

	type employee struct{
		name , project string
	}

	var m = make (map[string]string) // Map Syntax  = map[Keys]value

	m["A"] = "B" // insert in a map / it can be used to editing a value to the map also.
	m["C"] = "D"

	fmt.Println(m)

	var emp_map = make (map[string] employee)

	emp_map["emp1"] = employee{"N1" , "p1"}
	emp_map["emp2"] = employee{"N2", "p2"}

	fmt.Println(emp_map[""])

	val , ok := m["A"] // this is the syntax for checking if a value exists in the map or not
	emp_val , ok := emp_map["emp1"];
	fmt.Println(val , ok)
	fmt.Print(emp_val,ok)

}
