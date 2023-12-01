package main

import (
	"fmt"
	"strings"
)

type CurrencyInfo struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type CountryInfo struct {
	Name       struct{ Common, Official string } `json:"name"`
	Capital    []string                          `json:"capital"`
	Subregion  string                            `json:"subregion"`
	Currencies map[string]CurrencyInfo           `json:"currencies"`
}

const (
	currencySeparator = ", "
)

func (c CountryInfo) PrintInfo() {
	fmt.Printf("The capital of %s is: %s\nSubregion: %s\n", c.Name.Official, c.Capital[0], c.Subregion)

	var currencyList []string
	for _, currency := range c.Currencies {
		currencyList = append(currencyList, fmt.Sprintf("%s (%s)", currency.Name, currency.Symbol))
	}
	fmt.Println("Currency:", strings.Join(currencyList, currencySeparator))

	fmt.Println()
}
