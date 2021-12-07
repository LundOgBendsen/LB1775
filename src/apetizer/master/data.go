package main


// Import vs Include
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	// Unused imports are compile errors
	// "flag"
)


// Basic type definition
type IdType string

// Uppercase for public, lowercase for private
// Also showing inferred types
var Navn = "mctl"



// Structured data
type Tjeneste struct {
  // struct tags for reflection
	TjenesteId IdType `json: "name", xml: "navn"`
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
  // Anonymous variables
  // reg, _ := http.NewRequest()

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
	// Cross compile: GOOS=windows go build -o data.exe
	// https://mastedatabasen.dk/viskort/ContentPages/DataFraDatabasen.aspx
	// https://mastedatabasen.dk/Master/antenner/tjenester.json
	// https://mastedatabasen.dk/Master/antenner/teknologier.json
	// https://mastedatabasen.dk/Master/antenner.json?postnr=6900&tjeneste=2&teknologi=7&maxantal=15
	fmt.Println("Tjenester")
	
	// Range operator gives an interable (idx,e)
	for _, e := range getTjenester() {
		fmt.Printf("%v : %v\n", e.TjenesteId, e.Navn)
	}

  fmt.Println("Teknologier")
	for _, e := range getTeknologier() {
		fmt.Printf("%v : %v\n", e.Id, e.Navn)
	}
}
