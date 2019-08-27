package main

import (
	"errors"
	"log"
	"math"
)

func sqrt(val float64) (float64 , error){
	if val >0 {
		return  math.Sqrt(val) , nil
	}else{
		return 0 , errors.New("Math Error: value can not be negative for sqrt " )
	}
}

func main(){
	res , err := sqrt(10 )

	if err == nil{
		log.Println("Square root is : ",res)
	}else{
		log.Println(err)
	}
}
