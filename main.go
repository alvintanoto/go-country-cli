package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type CountryData []struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Tld         []string `json:"tld"`
	Cca2        string   `json:"cca2"`
	Ccn3        string   `json:"ccn3"`
	Cca3        string   `json:"cca3"`
	Independent bool     `json:"independent"`
	Status      string   `json:"status"`
	UnMember    bool     `json:"unMember"`
	Capital     []string `json:"capital"`
	Latlng      []int    `json:"latlng"`
	Area        float64  `json:"area"`
	Population  int      `json:"population"`
	Timezones   []string `json:"timezones"`
	Continents  []string `json:"continents"`
	CapitalInfo struct {
		Latlng []float64 `json:"latlng"`
	} `json:"capitalInfo"`
}

func main() {
	help := flag.Bool("help", false, "Show Help")
	find := flag.String("find",
		"",
		"Search by country name. It can be the native name or partial name")

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if len(*find) > 0 {
		findCountry(*find)
	}

}

func findCountry(country string) {
	resp, err := http.Get(fmt.Sprintf("https://restcountries.com/v3.1/name/%s", country))
	if err != nil {
		fmt.Println("Couldn't get country data")
		os.Exit(0)
	}

	if resp.StatusCode == 404 {
		fmt.Println("Country data not found")
		os.Exit(0)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Application error")
		os.Exit(0)
	}

	sb := string(body)

	var countries CountryData
	json.Unmarshal([]byte(sb), &countries)

	fmt.Println("")

	for _, val := range countries {
		fmt.Println("Official Name:", val.Name.Official)

		capitals := ""
		for i, capital := range val.Capital {
			capitals += capital

			if len(val.Capital)-1 != i {
				capitals += ", "
			}
		}
		fmt.Println("Capital City:", capitals)

		timezones := ""
		for i, timezone := range val.Timezones {
			timezones += timezone

			if len(val.Timezones)-1 != i {
				timezones += ", "
			}
		}
		fmt.Println("Timezones:", timezones)

		area := fmt.Sprintf("%.2f", val.Area)
		fmt.Println("Area:", area, "sqkm")
		fmt.Println("Population:", val.Population)

		countryCodes := fmt.Sprintf("%s, %s, %s", val.Cca2, val.Cca3, val.Ccn3)
		fmt.Println("Country Codes(CCA2, CCA3, CCN3):", countryCodes)
		fmt.Println("Independent:", val.Independent)
		fmt.Println("UN Member:", val.UnMember)

		fmt.Println("")
	}
}
