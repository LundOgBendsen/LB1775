package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Tjeneste struct {
	TjenesteId string `json:"id"`
	Navn       string
}

type Teknologi struct {
	Id   string
	Navn string
}

var Navn = "mflags"

func getJSON(url string, result interface{}) {
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

func visTjenester() {
	t := getTjenester()
	fmt.Println("Tjenester")
	for _, e := range t {
		fmt.Printf("%v(%v)\n", e.Navn, e.TjenesteId)
	}
}

func visTeknologier() {
	t := getTeknologier()
	fmt.Println("Teknologier")
	for _, e := range t {
		fmt.Printf("%v(%v)\n", e.Navn, e.Id)
	}
}

func main() {
	// https://mastedatabasen.dk/viskort/ContentPages/DataFraDatabasen.aspx
	// https://mastedatabasen.dk/Master/antenner.json?postnr=6900&tjeneste=2&teknologi=7&maxantal=15


  // Execute with subcommands having  different flags
  // tjeneste og teknologi viser de tilgængelige ting, med -h for hjælp
  // antenne tager diverse flag samt et postnummer
  
  // Default usage
  flag.Usage = func() {
    fmt.Fprintf(flag.CommandLine.Output(), Navn + " <-h|subcmd> \nWhere subcmd=<tjeneste|teknologi|antenne>\n")
  }


	// Command flags can operate in FlagSet
	tjenesteCmd := flag.NewFlagSet(Navn + " tjeneste", flag.ExitOnError)
	tjenesteCmd.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), Navn + " tjeneste - Viser de tilgængelige tjenester\n")
	}

	teknologiCmd := flag.NewFlagSet("teknologi", flag.ExitOnError)
  // 	teknologiCmd.Usage = func() {
  // 		fmt.Fprintf(flag.CommandLine.Output(), Navn + " teknologi - Viser de tilgængelige maste teknologier\n")
  // 	}
	
	
	listCmd := flag.NewFlagSet("antenne", flag.ExitOnError)
	listCmd.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), Navn + " antenne <postnummer> - Viser de tilgængelige antenner i angivne postnummer \n")
		listCmd.PrintDefaults()
	}
	listAntal := listCmd.Int("max", 15, "Max antal resultater")
	listTjeneste := listCmd.String("t", "", "Begræns til denne tjeneste type")
	listTeknologi := listCmd.String("T", "", "Begræns til denne masteteknologi")


  // Will parse the os.Args() for us
	flag.Parse()
	subCmd := flag.Args()

	if len(subCmd) < 1 {
	  flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "tjeneste":
		tjenesteCmd.Parse(subCmd[1:])
		visTjenester()

	case "teknologi":
		teknologiCmd.Parse(subCmd[1:])
		visTeknologier()

	case "antenne":
		listCmd.Parse(os.Args[2:])
		if len(listCmd.Args()) < 1 {
			fmt.Println("Angiv postnummer (eller -h)")
			os.Exit(1)
		}
		fmt.Println("Antal ", *listAntal)
		fmt.Println("Tjeneste ", *listTjeneste)
		fmt.Println("Teknologi ", *listTeknologi)
		fmt.Println("Postnummer ", listCmd.Arg(0))

	default:
	  flag.Usage()
		os.Exit(1)
	}

}
