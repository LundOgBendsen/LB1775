package main

import (
        "fmt"
        "github.com/mingrammer/cfmt"
)

func main() {
        fmt.Println("Hello world")
	cfmt.Infoln("For your information")
        cfmt.Successln("Success")
	cfmt.Warningln("This is a warning")
}
