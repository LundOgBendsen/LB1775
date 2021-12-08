package main


import (
	"fmt"
	"time"
	"math/rand"
)

func explainingChannels() {
  var value int
  
  // Declaring and initializing
  c := make(chan int)
  
  // Sending on a channel (block unless buffered channel)
  c <- 1
  
  // Receiving from a channel ("arrow"=data flow)
  value = <-c
  
  fmt.Println(value)
}


func explainingSelect() {
  c1 := make(chan string)
  c2 := make(chan string)
  c3 := make(chan int)
  
  select {
    case v1 := <- c1:
      fmt.Printf("Received ½v from c1\n", v1)
    case v2 := <- c2:
      fmt.Printf("Received ½v from c2\n", v2)
    case c3 <- 23:
      fmt.Printf("Sent %v to c3\n", 23)
    default:
      fmt.Printf("No channel was ready to communicate")
  }
}



func generator(msg string) <-chan string { // Return recieve-only channel of strings
  c:= make(chan string)
  go func() { // We launch the goroutine from inside the function (anonymously)
    for i := 0 ; ; i++ {
      c <- fmt.Sprintf("%s %d",msg, i)
      time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    }
  }()
  return c // Return the channel to the caller
}


// Fan-in of multiple channels
func fanIn(input1, input2 <-chan string) <-chan string {
  c:= make(chan string)
  go func() { for { c <- <-input1 } }()
  go func() { for { c <- <-input2 } }()
  return c
}


// Fan-in of multiple channels using select
func fanInSelect(input1, input2 <-chan string) <-chan string {
  c:= make(chan string)
  go func() {
    for {
      select {
        case s:= <- input1: c <-s
        case s:= <- input2: c <-s
      }
    }
  }()
  return c
}

// Either receive or timeout
func receiveOrTimeout() {
  c:=generator("Joe")
  for {
    select {
      case s:= <- c: fmt.Println(s)
      case <- time.After(1 * time.Second):
        fmt.Println("You're too slow")
        return
    }
  }
}


// Push advanced data into channel with return channel
type Message struct {
  str string
  wait chan bool // Send this channel for acknowledgement
}




func main() {
  c:= generator("Hello world!") // FUnction returning a channel
  for i:= 0 ; i < 5 ; i++ {
    fmt.Println(<-c)
  }
  
  
  mike := generator("mike")
  dave := generator("dave")
  for i:= 0 ; i < 5 ; i++ {
    fmt.Println(<-mike)
    fmt.Println(<-dave)
  }
  
  fanin := fanIn(mike,dave)
  for i:= 0 ; i < 5 ; i++ {
    fmt.Println(<-fanin)
  }
  
  
  fmt.Println("Goodbye")
}
