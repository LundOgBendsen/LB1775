package main

import (
	"fmt"
	"strings"
)

func main(){
	var input string
	fmt.Println("Enter string")
	fmt.Scan(&input)
	input = strings.ToLower(input)
	if strings.Contains(input, "wally") {
		fmt.Println("Found!")
	}else {
		fmt.Println("NotFound!")
	}
}
