package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	// Include in version latest (since nothing else is specified e.g. @v1.1.0)
	"github.com/mingrammer/cfmt"
)

type Tjeneste struct {
	TjenesteId string `json:"id"`
	Navn       string
}

type Teknologi struct {
	Id   string
	Navn string
}

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
	cfmt.Successln("Tjenester")
	for _, e := range t {
		cfmt.Infof("%v(%v)\n", e.Navn, e.TjenesteId)
	}
}

func visTeknologier() {
	t := getTeknologier()
	cfmt.Successln("Teknologier")
	for _, e := range t {
		cfmt.Infof("%v(%v)\n", e.Navn, e.Id)
	}
}


type ListRequest struct {
  antal int
  tjeneste string
  teknologi string
  postnummer string
}


func (request ListRequest) postnummerURI() string {
  return "?postnr="+request.postnummer
}

func (request ListRequest) tjenesteURI() (id string) {
  if request.tjeneste != "" {
    t:= getTjenester()
    for _, e := range t {
      if (e.Navn == request.tjeneste) {
        id = e.TjenesteId
      }
    }
    if id != "" {
      id = "&tjeneste=" + id
    }
  }
  return
}

func (request ListRequest) teknologiURI() (id string) {
  if request.teknologi != "" {
    t:= getTeknologier()
    for _, e := range t {
      if (e.Navn == request.teknologi) {
        id = e.Id
      }
    }
    if id != "" {
      id = "&teknologi=" + id
    }
  }
  return
}

func (request ListRequest) antalURI() string {
  return "&maxantal=" + strconv.Itoa(request.antal)
}

func (request ListRequest) toURL() string {
  return "https://mastedatabasen.dk/Master/antenner.json" + request.postnummerURI() + request.tjenesteURI() + request.teknologiURI() + request.antalURI()
}

type Antenne struct {
  Vejnavn VejnavnStruct
  Husnr string
  Idriftsaettelsesdato string
  ForventetIdriftsaettelsesdato string `json:"forventet_idriftsaettelsesdato"`
  TjenesteArt Tjeneste
  Teknologi Teknologi
  Frekvensbaand string
}

type VejnavnStruct struct {
  Kode string
  Navn string
}

func visMaster(antal int, tjeneste string, teknologi string, postnummer string) {
  request := ListRequest { antal:antal, tjeneste:tjeneste, teknologi:teknologi, postnummer:postnummer }
  url := request.toURL()
  cfmt.Successln(url)

  var antenner []Antenne
  getJSON(url, &antenner)
  
  for idx,a := range antenner {
    addr := a.Vejnavn.Navn + " " + a.Husnr
    freq := a.Frekvensbaand
    if freq != "" {
      freq = freq + "MHz"
    }
    cfmt.Infof("%03v: %-30v  %v  %v  %8v  %8v\n", idx, addr, a.Idriftsaettelsesdato, a.TjenesteArt.Navn, a.Teknologi.Navn, freq)
  }
}

func main() {
	// https://mastedatabasen.dk/viskort/ContentPages/DataFraDatabasen.aspx

	listCmd := flag.NewFlagSet("antenne", flag.ExitOnError)
	listCmd.Usage = func() {
		cfmt.Fwarningf(flag.CommandLine.Output(), "mctl antenne <postnummer> - Viser de tilgængelige antenner i angivne postnummer \n")
		listCmd.PrintDefaults()
	}
	listAntal := listCmd.Int("max", 15, "Max antal resultater")
	listTjeneste := listCmd.String("t", "", "Begræns til denne tjeneste type")
	listTeknologi := listCmd.String("T", "", "Begræns til denne masteteknologi")

	tjenesteCmd := flag.NewFlagSet("mctl tjeneste", flag.ExitOnError)
	tjenesteCmd.Usage = func() {
		cfmt.Fwarningf(flag.CommandLine.Output(), "mctl tjeneste - Viser de tilgængelige tjenester\n")
	}

	teknologiCmd := flag.NewFlagSet("teknologi", flag.ExitOnError)
	teknologiCmd.Usage = func() {
		cfmt.Fwarningf(flag.CommandLine.Output(), "mctl teknologi - Viser de tilgængelige maste teknologier\n")
	}

	if len(os.Args) < 2 {
		cfmt.Errorln("Angiv underkommando (tjeneste, teknologi, antenne)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "tjeneste":
		tjenesteCmd.Parse(os.Args[2:])
		visTjenester()

	case "teknologi":
		teknologiCmd.Parse(os.Args[2:])
		visTeknologier()

	case "antenne":
		listCmd.Parse(os.Args[2:])
		if len(listCmd.Args()) < 1 {
			cfmt.Warningln("Angiv postnummer (eller -h)")
			os.Exit(1)
		}
		visMaster(*listAntal, *listTjeneste, *listTeknologi, listCmd.Arg(0))

	default:
		cfmt.Warningln("Forventede <tjeneste|teknologi|antenne>")
		os.Exit(1)
	}

}
