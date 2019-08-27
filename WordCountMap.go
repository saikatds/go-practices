package main

import (
	"fmt"
)

func wordCount( word string) map[string]int{

	var count_map = make (map[string]int)

	for i:= 0 ; i < len(word) ; i++{
		char := string(word[i])

		val , ok := count_map[char]

		if ok == true {
			count_map[char] = val+1
		}else{
			count_map[char] = 1
		}
	}

	return count_map

}

func main(){
	word := "saikat dey sarkar"
	fmt.Println(wordCount(word))
}