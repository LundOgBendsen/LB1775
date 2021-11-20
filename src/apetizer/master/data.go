package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Structured data with 'json-annotations'
type Tjeneste struct {
	TjenesteId string `json:"id"`
	Navn       string
}

type Teknologi struct {
  Id string
  Navn string
}

func getJSON(url string, result interface{})  {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "go-cmd-line")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, result)
  
}


func getTjenester() []Tjeneste {
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
	// Cross compile: GOOS=windows go build -o hello-world.exe
	// https://mastedatabasen.dk/viskort/ContentPages/DataFraDatabasen.aspx
	// https://mastedatabasen.dk/Master/antenner/tjenester.json
	// https://mastedatabasen.dk/Master/antenner/teknologier.json
	// https://mastedatabasen.dk/Master/antenner.json?postnr=6900&tjeneste=2&teknologi=7&maxantal=15
	fmt.Println("Tjenester")
	for _, e := range getTjenester() {
		fmt.Printf("%v : %v\n", e.TjenesteId, e.Navn)
	}

  fmt.Println("Teknologier")
	for _, e := range getTeknologier() {
		fmt.Printf("%v : %v\n", e.Id, e.Navn)
	}
}
