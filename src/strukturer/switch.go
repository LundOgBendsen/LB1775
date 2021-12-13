package main

import "fmt"

func main() {

  x := 5
  y := 12

  switch {
    case x > 10: fmt.Println("First case")
    case y > 10: fmt.Println("Second case")
  }
}
