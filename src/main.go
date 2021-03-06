package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var stocks []Stock
	optSymbolPtr := flag.String("symbol", "", "The stock symbol to lookup. Can pass multiple comma-delimited symbols.")
	optJSONPtr := flag.Bool("json", false, "Output will be JSON formatted.")
	//optFieldsPtr := flag.String("fields", "", "Comma separated list of fields to select in the output")
	flag.Parse()
	if *optSymbolPtr == "" {
		fmt.Println("Error: no arguments provided.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	//fieldsArgs := strings.Split(*optFieldsPtr, ",")
	symbolArgs := strings.Split(*optSymbolPtr, ",")
	stocks, err := GetQuotesBulk(symbolArgs)
	if *optJSONPtr == true {
		// Convert output into JSON
		outJSON, e := json.Marshal(stocks)
		if e != nil {
			fmt.Printf("Error during JSON conversion: %v", e)
			os.Exit(1)
		}
		fmt.Println(string(outJSON))
		os.Exit(0)
	}
	if err != nil {
		fmt.Printf("Error attempting to fetch stocks: %v\n", err)
		os.Exit(1)
	}
	GenTable(&stocks)
}
