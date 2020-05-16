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

// Stock struct of a single ticker/symbol containing
// various information about said ticker.
type Stock struct {
	TradingHalted                  bool
	Low52Weeks                     float64
	High52Weeks                    float64
	AskSize                        int
	Ask                            float64
	BidSize                        int
	Bid                            float64
	PercentChangeFromPreviousClose float64
	ChangeFromPreviousClose        float64
	PreviousClose                  float64
	Volume                         int
	LastSize                       int
	Last                           float64
	Low                            float64
	High                           float64
	Close                          float64
	Open                           float64
	UTCOffset                      int
	Delay                          float64
	Outcome                        string
	Security                       struct {
		MostLiquidExchange       bool
		CategoryOrIndustry       string
		MarketIdentificationCode string
		Market                   string
		Name                     string
		Valoren                  string
		ISIN                     interface{}
		Symbol                   string
		CUSIP                    interface{}
		CIK                      string
	}
	IdentifierType               string
	Identifier                   string
	LastMarketIdentificationCode string
	AskMarketIdentificationCode  string
	BidMarketIdentificationCode  string
	Currency                     string
	AskTime                      string
	AskDate                      string
	BidTime                      string
	BidDate                      string
	PreviousCloseDate            string
	Time                         string
	Date                         string
	Identity                     string
	Message                      string
}

func isValidSymbol(symbol string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9.:]+$`)
	return re.MatchString(symbol)
}

// GetQuotesBulk given a string slice, returns
// a slice of Stock structs containing requested
// stock information.
// See: GetQuotes
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

// GetQuote returns a Stock struct
// for a single stock ticker that is passed
// in as a Stock struct
func GetQuote(symbol Stock) (Stock, error) {
	if !isValidSymbol(symbol.Identifier) {
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
