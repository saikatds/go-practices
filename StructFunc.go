package main

import (
	"fmt"
)

type employee struct {
	name string
	dept string
	empId int
	project string
}

/**
This is a normal function that we can call
 */
func isWorking( emp employee) bool{
	if emp.project == "" {
		return false
	}else {
		return true
	}
}

/**
This function will work as a method to the struct employee
 */
func (emp *employee )allocateWork(projectName string){
	fmt.Println("Allocating employee to ",projectName)
	emp.project = projectName
}

/**
This function will work as a method to the struct employee
 */
func (emp *employee)deallocateWork(){
	fmt.Println("De allocating employee from ",emp.project)
	emp.project = ""
}

func main(){
	emp1:= employee{"A","MicroService",10000 ,  ""}
	fmt.Println("Employee Details : ",emp1.name,emp1.project)
	fmt.Println("Employee working : ",isWorking(emp1))
	emp1.allocateWork("mast")
	fmt.Println("Employee Details : ",emp1.name,emp1.project)
	fmt.Println("Employee working : ",isWorking(emp1))

}



