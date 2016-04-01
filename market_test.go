package main

import "testing"

func TestStubMarketStocks(t *testing.T) {
	var market Market
	stubMarketStocks(&market)

	numberOfStocks := len(market.Stocks)
	if numberOfStocks != 3 {
		t.Errorf("Was expecting 3 stocks in the market, got %v", numberOfStocks)
	}
}

func TestInitializeMarket(t *testing.T) {
	market, err := initializeMarket()

	if (err != nil) {
		t.Errorf("Got an error creating the market: %v", err)
	}

	if &market == nil {
		t.Errorf("InitializeMarket returned nil")
	}

	if len(market.Orders) != 0 {
		t.Errorf("Market orders didn't get initialized.")
	}
}

