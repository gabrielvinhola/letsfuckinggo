package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	apiEndpoint   = "https://restcountries.com/v3.1/all?fields=name,capital,currencies,subregion"
	successStatus = http.StatusOK
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	responseData, err := fetchData(apiEndpoint)
	checkError(err)

	var countries []CountryInfo
	err = json.Unmarshal(responseData, &countries)
	checkError(err)

	userInputs := ReadUserInput()

	for _, userInputCountry := range userInputs {
		found := false
		for _, country := range countries {
			if strings.EqualFold(country.Name.Common, userInputCountry) {
				country.PrintInfo()
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Country %s not found in the data\n", userInputCountry)
		}
	}
}
