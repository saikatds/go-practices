package main

import (
	"encoding/json"
	"log"
)

type response1 struct {
	page int
	fruits []string
}

type response2 struct {
	page int `json:"page"`
	fruits []string `json:"fruits"`
}

func main(){
	boolVar , error := json.Marshal(false)

	if error != nil{
		log.Fatal(error)
	}else{
		log.Println(string(boolVar))
	}

	strB, _ := json.Marshal("gopher")
	log.Println(string(strB))
}
