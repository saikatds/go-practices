package main

import "fmt"

func main(){
	arr := make ([] int , 10)

	for i := range  arr{
		arr[i] =  i+3
	}

	/**
		range gives index and values
	 */
	for index,val := range arr{
		fmt.Println("Value : ",val , " index : ",index)
	}

}
