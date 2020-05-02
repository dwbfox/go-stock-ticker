package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var stocks []Stock

	// Parse argv and append to stock list
	cmdargs := os.Args[1:]

	if len(cmdargs) <= 0 {
		fmt.Println("Please provide a stock symbol to quote.")
		os.Exit(1)
	}

	for _, opt := range cmdargs {
		symbol := Stock{
			Identifier: strings.ToUpper(opt),
		}
		stocks = append(stocks, symbol)
	}

	for _, symbol := range stocks {
		quote, err := GetQuote(symbol)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		fmt.Printf(">> Stock quote for %s: $%.2f\n", quote.Identifier, quote.Ask)
	}
}
