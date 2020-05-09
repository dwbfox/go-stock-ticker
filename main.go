package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var stocks []Stock
	symbolPtr := flag.String("symbol", "", "The stock symbol to lookup. Can pass multiple comma-delimited symbols.")
	flag.Parse()
	symbolArgs := strings.Split(*symbolPtr, ",")
	stocks, err := GetQuotesBulk(symbolArgs)
	if err != nil {
		panic(fmt.Sprintf("Error attempting to fetch stocks: %v", err))
	}
	GenTable(&stocks)
}
