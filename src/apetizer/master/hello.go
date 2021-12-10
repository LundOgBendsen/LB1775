package main


// Import vs Include
import (
	"fmt"
	// Unused imports are compile errors
	// "flag"
)


// Uppercase for public, lowercase for private
var Navn string = "audience"

// Inferred type
//var Navn = "audience"


func greeting(name string) string {
  return "Hello, " + name
}

func main() {
	// Cross compile: GOOS=windows go build -o data.exe
	fmt.Println(greeting(Navn))
}
