package main


import (
	"fmt"
	"time"
	"math/rand"
)



var (
  Web = fakeSearch("web")
  Image = fakeSearch("image")
  Video = fakeSearch("video")
)

// For the replicated stuff
var (
  Web1 = fakeSearch("web")
  Image1 = fakeSearch("image")
  Video1 = fakeSearch("video")
  Web2 = fakeSearch("web")
  Image2 = fakeSearch("image")
  Video2 = fakeSearch("video")
)

type Search func(query string) Result
type Result string


func fakeSearch(kind string) Search {
  return func(query string) Result {
    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    return Result(fmt.Sprintf("%s result for %s\n", kind, query))
  }
}


// Primitive, sequential
func Google(query string) (results []Result) {
  results = append(results, Web(query))
  results = append(results, Image(query))
  results = append(results, Video(query))
  return
}

// First attempt of parallel
func Google2(query string) (results []Result) {
  c := make(chan Result)
  go func() { c <- Web(query) }()
  go func() { c <- Image(query) }()
  go func() { c <- Video(query) }()
  
  for i:= 0 ; i < 3; i++ {
    results = append(results, <-c)
  }
  return
}

// First attempt of parallel but with timeout
func Google2point1(query string) (results []Result) {
  c := make(chan Result)
  go func() { c <- Web(query) }()
  go func() { c <- Image(query) }()
  go func() { c <- Video(query) }()
  
  for i:= 0 ; i < 3; i++ {
    select {
      case result := <- c: results = append(results, result)
      case <-time.After(80 * time.Millisecond):
        fmt.Println("Timeout ")
        return
    }
  }
  return
}

func First(query string, replicas ...Search) Result {
  c := make(chan Result)
  searchReplica := func(i int) { c <- replicas[i](query) }
  for i := range replicas {
    // Run each in background, and let the first one answer first..
    go searchReplica(i)
  }
  return <-c // Since we just need one read...
}


// Parallel replicated, reducing tail latency using replicated search servers
func Google3(query string) (results []Result) {
  c := make(chan Result)
  go func() { c <- First(query, Web1, Web2) }()
  go func() { c <- First(query, Image1, Image2) }()
  go func() { c <- First(query, Video1, Video2) }()
  
  for i:= 0 ; i < 3; i++ {
    results = append(results, <-c)
  }
  return
}


// Turn a slow, sequential, failure-sensitive search into a fast, concurrent, replicated, robust search
func main() {
  rand.Seed(time.Now().UnixNano())
  start := time.Now()
  //results := Google("golang")
  //results := Google2point("golang")
  //results := Google2("golang")
  results := Google3("golang")
  elapsed := time.Since(start)
  fmt.Println(results)
  fmt.Println(elapsed)
}
