package main

import (
	"testing"

	"github.com/dob/marketsim"
)

func TestStubMarketStocks(t *testing.T) {
	var market *marketsim.Market = marketsim.NewMarket()
	stubMarketStocks(market)

	numberOfStocks := len(market.Stocks)
	if numberOfStocks <= 1 {
		t.Errorf("Was expecting a bunch of stocks in the market, got %v", numberOfStocks)
	}
}

func TestInitializeMarketWithStocks(t *testing.T) {
	market, err := initializeMarketWithStocks()

	if err != nil {
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
	market := marketsim.NewMarket()
	stubMarketStocks(market)
	syms := market.Symbols()

	if len(syms) <= 1 {
		t.Errorf("The symbols extraction only extracted %v symbols", len(syms))
	}
}

func TestOrderPrintingBuy(t *testing.T) {
	o := marketsim.Order{"AMZN", marketsim.BuyOrderType, marketsim.LimitOrderType, 20, 45, marketsim.OrderStatusOpen}
	if o.String() != "Buy: 20 shares of AMZN at $45." {
		t.Errorf("Got a bad buy string: %v.", o.String())
	}
}

func TestOrderPrintingSell(t *testing.T) {
	o := marketsim.Order{"AMZN", marketsim.SellOrderType, marketsim.LimitOrderType, 20, 45, marketsim.OrderStatusOpen}
	if o.String() != "Sell: 20 shares of AMZN at $45." {
		t.Errorf("Got a bad sell string: %v", o.String())
	}
}

func TestOrderTypeString(t *testing.T) {
	o := marketsim.Order{"AMZN", marketsim.SellOrderType, marketsim.LimitOrderType, 20, 45, marketsim.OrderStatusOpen}
	if o.OrderType.String() != "Limit" {
		t.Errorf("Got a bad order type string")
	}

	o = marketsim.Order{"AMZN", marketsim.SellOrderType, marketsim.MarketOrderType, 20, 45, marketsim.OrderStatusOpen}
	if o.OrderType.String() != "Market" {
		t.Errorf("Got a bad order type string")
	}
}
