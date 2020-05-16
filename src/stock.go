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
	TradingHalted                  bool    `json:"TradingHalted"`
	Low52Weeks                     float64 `json:"Low52Weeks"`
	High52Weeks                    float64 `json:"High52Weeks"`
	AskSize                        int     `json:"AskSize"`
	Ask                            float64 `json:"Ask"`
	BidSize                        int     `json:"BidSize"`
	Bid                            float64 `json:"Bid"`
	PercentChangeFromPreviousClose float64 `json:"PercentChangeFromPreviousClose"`
	ChangeFromPreviousClose        float64 `json:"ChangeFromPreviousClose"`
	PreviousClose                  float64 `json:"PreviousClose"`
	Volume                         int     `json:"Volume"`
	LastSize                       int     `json:"LastSize"`
	Last                           float64 `json:"Last"`
	Low                            float64 `json:"Low"`
	High                           float64 `json:"High"`
	Close                          float64 `json:"Close"`
	Open                           float64 `json:"Open"`
	UTCOffset                      int     `json:"UTCOffset"`
	Delay                          float64 `json:"Delay"`
	Outcome                        string  `json:"Outcome"`
	Security                       struct {
		MostLiquidExchange       bool        `json:"MostLiquidExchange"`
		CategoryOrIndustry       string      `json:"CategoryOrIndustry"`
		MarketIdentificationCode string      `json:"MarketIdentificationCode"`
		Market                   string      `json:"Market"`
		Name                     string      `json:"Name"`
		Valoren                  string      `json:"Valoren"`
		ISIN                     interface{} `json:"ISIN"`
		Symbol                   string      `json:"Symbol"`
		CUSIP                    interface{} `json:"CUSIP"`
		CIK                      string      `json:"CIK"`
	} `json:"Security"`
	IdentifierType               string `json:"IdentifierType"`
	Identifier                   string `json:"Identifier"`
	LastMarketIdentificationCode string `json:"LastMarketIdentificationCode"`
	AskMarketIdentificationCode  string `json:"AskMarketIdentificationCode"`
	BidMarketIdentificationCode  string `json:"BidMarketIdentificationCode"`
	Currency                     string `json:"Currency"`
	AskTime                      string `json:"AskTime"`
	AskDate                      string `json:"AskDate"`
	BidTime                      string `json:"BidTime"`
	BidDate                      string `json:"BidDate"`
	PreviousCloseDate            string `json:"PreviousCloseDate"`
	Time                         string `json:"Time"`
	Date                         string `json:"Date"`
	Identity                     string `json:"Identity"`
	Message                      string `json:"Message"`
}

func IsValidSymbol(symbol string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9.:]+$`)
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
	//log.Println(rawResult)
	json.Unmarshal([]byte(rawResult), &symbol)
	if symbol.Outcome != "Success" {
		return Stock{}, fmt.Errorf("api call succeeded, but symbol was invalid: %s", symbol.Identifier)
	}
	return symbol, nil
}

func GenTable(stocks *[]Stock) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Symbol", "Name", "Industry", "Ask", "Bid", "Volume"})
	for _, symbol := range *stocks {
		quote, err := GetQuote(symbol)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		table.Append([]string{
			quote.Identifier,
			quote.Security.Name,
			quote.Security.CategoryOrIndustry,
			fmt.Sprintf("$%0.2f", quote.Ask),
			fmt.Sprintf("$%0.2f", quote.Bid),
			fmt.Sprintf("%d", quote.Volume),
		})
	}
	table.Render()
	return nil
}
