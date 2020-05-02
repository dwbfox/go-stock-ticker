package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Stock struct {
	PreviousClose float32
	Name          string
	Volume        int64
	Identifier    string
	Ask           float32
	Bid           float32
}

func GetQuote(stock Stock) (Stock, error) {
	endpoint := "https://duckduckgo.com/js/spice/stocks/%s"
	resp, err := http.Get(fmt.Sprintf(endpoint, stock.Identifier))
	if err != nil {
		fmt.Println("Failed to fetch data for", stock.Identifier, ": ", err)
		return stock, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return stock, err
	}
	rawResult := string(body)
	rawResult = strings.ReplaceAll(rawResult, "ddg_spice_stocks(", "")
	rawResult = strings.ReplaceAll(rawResult, ");", "")
	// parse the clenaed up JSON
	json.Unmarshal([]byte(rawResult), &stock)
	return stock, nil
}
