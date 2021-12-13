package main

import (
	"fmt"
)

func main() {
	var f float64

	fmt.Println("Enter a floating point number")
	fmt.Scan(&f)

	i := int(f)

	fmt.Printf("The integer portion is %v\n", i)
}
