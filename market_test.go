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

func TestMarketSymbols(t *testing.T) {
	var market Market
	stubMarketStocks(&market)
	syms := market.Symbols()

	if len(syms) != 3 {
		t.Errorf("The symbols extraction only extracted %v symbols", len(syms))
	}
}

func TestOrderPrintingBuy(t *testing.T) {
	o := Order{"AMZN", BuyOrderType, 20, 45}
	if o.String() != "Buy: 20 shares of AMZN at $45" {
		t.Errorf("Got a bad buy string: %v", o.String())
	}
}

func TestOrderPrintingSell(t *testing.T) {
	o := Order{"AMZN", SellOrderType, 20, 45}
	if o.String() != "Sell: 20 shares of AMZN at $45" {
		t.Errorf("Got a bad sell string: %v", o.String())
	}
}
