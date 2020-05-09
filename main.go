package main

import (
	"fmt"
	"os"
)

func main() {
	var stocks []Stock
	cmdargs := os.Args[1:]
	if len(cmdargs) <= 0 {
		fmt.Println("Please provide a stock symbol to quote.")
		os.Exit(1)
	}
	stocks, err := GetQuotesBulk(cmdargs)
	if err != nil {
		panic(fmt.Sprintf("Error attempting to fetch stocks: %v", err))
	}
	GenTable(&stocks)
}
