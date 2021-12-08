package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


var Navn = "mctl"

// Basic type definition
type IdType string



// Structured data
type Tjeneste struct {
  // struct tags for reflection (annotations)
	TjenesteId IdType `json:"id", xml:"id"`
	Navn       string
}

type Teknologi struct {
  Id IdType
  Navn string
}

func getJSON(url string, result interface{})  {

  // Variable initialization
  // var client http.Client
  // var client = http.Client()

	client := http.Client{
		Timeout: time.Second * 2,
	}

  // multiple returns return a,b,c,d,e,f,g
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "go-cmd-line")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

  // Deferred execution (But not resolution)
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

  // Json marshal into all public fields using optional struct tags
	json.Unmarshal(body, result)
  
}


func getTjenester() []Tjeneste {

  // Slices vs Arrays
  var t []Tjeneste
	getJSON("https://mastedatabasen.dk/Master/antenner/tjenester.json", &t)
	return t
}

func getTeknologier() []Teknologi {
  var t []Teknologi
  getJSON("https://mastedatabasen.dk/Master/antenner/teknologier.json", &t)
  return t
}

func main() {
	// https://mastedatabasen.dk/viskort/ContentPages/DataFraDatabasen.aspx
	// https://mastedatabasen.dk/Master/antenner/tjenester.json
	// https://mastedatabasen.dk/Master/antenner/teknologier.json
	fmt.Println("Tjenester")
	
	// Range operator gives an interable (idx,e)
	// Anonymous variables (otherwise compile error)
	for _, e := range getTjenester() {
		fmt.Printf("%v : %v\n", e.TjenesteId, e.Navn)
	}

  fmt.Println("Teknologier")
	for _, e := range getTeknologier() {
		fmt.Printf("%v : %v\n", e.Id, e.Navn)
	}
}
