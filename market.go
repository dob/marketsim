package main

import (
	"fmt"
	"bytes"
)

// Equity representing one stock
type Stock struct {
	Symbol string
	Name string
	Bid float64
	Offer float64
}

func (s Stock) String() string {
	return fmt.Sprintf("%v, %v: %v-%v", s.Symbol, s.Name, s.Bid, s.Offer)
}

type OrderTypeVal int
const (
	BuyOrderType OrderTypeVal = iota
	SellOrderType OrderTypeVal = iota
)

// Represents an order to buy or sell
type Order struct {
	Symbol string
	OrderType OrderTypeVal
	Shares int
	Value float64
}

// Establish the market
type Market struct {
	Stocks map[string]Stock
	Orders []Order
}

func (m Market) String() string {
	var marketOutput bytes.Buffer

	marketOutput.WriteString("Current Market Status\n")
	marketOutput.WriteString("=====================\n")

	for _, stock := range m.Stocks {
		marketOutput.WriteString(stock.String() + "\n")
	}

	return marketOutput.String()
}

// Populate the market with some fake data
func stubMarketStocks(m *Market) {
	m.Stocks = make(map[string]Stock)

	sym1 := Stock{"AMZN", "Amazon", 568.2, 568.4}
	m.Stocks[sym1.Symbol] = sym1

	sym2 := Stock{"TSLA", "Tesla Motors", 288.5, 289}
	m.Stocks[sym2.Symbol] = sym2

	sym3 := Stock{"TWTR", "Twitter", 15.4, 15.60}
	m.Stocks[sym3.Symbol] = sym3
}

// Initialize the market
func initializeMarket() (Market, error) {
	var market Market
	stubMarketStocks(&market)
	market.Orders = make([]Order, 0)

	return market, nil
}

func main() {
	market, _ := initializeMarket()

	fmt.Println(market)
}
