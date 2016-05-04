package main

import (
	"testing"

	"github.com/dob/marketsim/shared/datatypes"
)

func TestStubMarketStocks(t *testing.T) {
	var market *dt.Market = dt.NewMarket()
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
	market := dt.NewMarket()
	stubMarketStocks(market)
	syms := market.Symbols()

	if len(syms) <= 1 {
		t.Errorf("The symbols extraction only extracted %v symbols", len(syms))
	}
}

func TestOrderPrintingBuy(t *testing.T) {
	o := dt.Order{"AMZN", dt.BuyOrderType, dt.LimitOrderType, 20, 45, dt.OrderStatusOpen}
	if o.String() != "Buy: 20 shares of AMZN at $45." {
		t.Errorf("Got a bad buy string: %v.", o.String())
	}
}

func TestOrderPrintingSell(t *testing.T) {
	o := dt.Order{"AMZN", dt.SellOrderType, dt.LimitOrderType, 20, 45, dt.OrderStatusOpen}
	if o.String() != "Sell: 20 shares of AMZN at $45." {
		t.Errorf("Got a bad sell string: %v", o.String())
	}
}

func TestOrderTypeString(t *testing.T) {
	o := dt.Order{"AMZN", dt.SellOrderType, dt.LimitOrderType, 20, 45, dt.OrderStatusOpen}
	if o.OrderType.String() != "Limit" {
		t.Errorf("Got a bad order type string")
	}

	o = dt.Order{"AMZN", dt.SellOrderType, dt.MarketOrderType, 20, 45, dt.OrderStatusOpen}
	if o.OrderType.String() != "Market" {
		t.Errorf("Got a bad order type string")
	}
}
