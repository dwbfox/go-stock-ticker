package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var stocks []Stock
	symbolPtr := flag.String("symbol", "", "The stock symbol to lookup. Can pass multiple comma-delimited symbols.")
	flag.Parse()
	if *symbolPtr == "" {
		fmt.Println("Error: no arguments provided.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	symbolArgs := strings.Split(*symbolPtr, ",")
	stocks, err := GetQuotesBulk(symbolArgs)
	if err != nil {
		fmt.Printf("Error attempting to fetch stocks: %v", err)
		os.Exit(1)
	}
	GenTable(&stocks)
}
