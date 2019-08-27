package main

import "fmt"

func main(){
	/* go chanel */
	messages := make(chan string)

	/* go routine */
	go func() {
		messages <- "ping"
	}()

	msg := <- messages

	fmt.Println(msg)
}
