package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Stock struct {
	PreviousClose float32
	Volume        int64
	Identifier    string
	Ask           float32
	Bid           float32
	Message       string
	Outcome       string
	AskSize       uint32
	BidSize       uint32
	Time          string
	Date          string
	Security      string
}

func IsValidSymbol(symbol string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9.]+$`)
	return re.MatchString(symbol)
}

func GetQuotesBulk(symbols []string) ([]Stock, error) {
	var stocks []Stock
	for _, symbol := range symbols {
		stock, err := GetQuote(Stock{
			Identifier: symbol,
		})
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

func GetQuote(symbol Stock) (Stock, error) {
	if !IsValidSymbol(symbol.Identifier) {
		return Stock{}, fmt.Errorf("invalid stock symbol provided: %s", symbol.Identifier)
	}
	endpoint := "https://duckduckgo.com/js/spice/stocks/%s"
	resp, err := http.Get(fmt.Sprintf(endpoint, symbol.Identifier))
	if err != nil {
		fmt.Println("Failed to fetch data for", symbol.Identifier, ": ", err)
		return symbol, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return symbol, err
	}
	rawResult := string(body)
	rawResult = strings.ReplaceAll(rawResult, "ddg_spice_stocks(", "")
	rawResult = strings.ReplaceAll(rawResult, ");", "")
	json.Unmarshal([]byte(rawResult), &symbol)
	if symbol.Outcome != "Success" {
		return Stock{}, fmt.Errorf("api call succeeded, but symbol was invalid: %s", symbol.Identifier)
	}
	return symbol, nil
}

func GenTable(stocks *[]Stock) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Symbol", "Ask", "Bid", "Volume"})
	for _, symbol := range *stocks {
		quote, err := GetQuote(symbol)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		table.Append([]string{
			quote.Identifier,
			fmt.Sprintf("$%0.2f", quote.Ask),
			fmt.Sprintf("$%0.2f", quote.Bid),
			fmt.Sprintf("%d", quote.Volume),
		})
	}
	table.Render()
	return nil
}
