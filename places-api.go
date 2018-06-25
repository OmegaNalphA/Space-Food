package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var placesKey = "nah"

type placeid struct {
	PlaceID string `json:"place_id"`
}
type place struct {
	Candidates []placeid `json:"candidates"`
}

type person struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

type people struct {
	People []person `json:"people"`
	// People []map[string]interface{} `json:"people"`
}

var pln = fmt.Println

func main() {
	// peopleInSpace()
	places("Museum of Contemporary Art Australia")
}

func places(inputText string) {
	url := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json"

	placesClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Testing how to query API's from parameters")
	q := req.URL.Query()
	q.Add("key", placesKey)
	q.Add("input", inputText)
	q.Add("inputtype", "textquery")
	req.URL.RawQuery = q.Encode()
	// pln(req.URL.String())

	res, getErr := placesClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	// pln(res.StatusCode)
	if res.StatusCode == 404 {
		log.Fatal("Hit a 404")
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// n := len(body)
	// s := string(bod0y)
	// pln(s)

	output := place{}
	jsonErr := json.Unmarshal(body, &output)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	pln(output)
}

func peopleInSpace() {
	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
	// textBytes := []byte(text)

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for _, p := range people1.People {
		pln(p.Name)
	}
}
