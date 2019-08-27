package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter , r *http.Request){
	fmt.Fprint(w,"First Go web application",r)
}

func aboutHandler(w http.ResponseWriter , r *http.Request){
	fmt.Fprint(w,"About section of the go web application")
}

func main(){

	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/about",aboutHandler)
	http.ListenAndServe(":8000" , nil)
}
