package main

import (
	"log"
	"testing"
)

func TestInvalidStock(t *testing.T) {
	invalidStock := Stock{}
	result, err := GetQuote(invalidStock)
	if err == nil && result.Identifier == "0" {
		t.Errorf("Got GOOD result for known BAD stock struct: %v\n", err)
		t.Failed()
	}
}

func TestValidStock(t *testing.T) {
	goodStock := Stock{
		Identifier: "SPY",
	}

	result, err := GetQuote(goodStock)
	if err != nil {
		t.Errorf("Error result for known good stock: %v\n", err)
	}

	if result.PreviousClose == 0.00 {
		t.Errorf("Unexpected result from known good stock: %v\n", result)
	}

	log.Printf("Resulting stock struct: %v", result)
}
