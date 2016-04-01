package main

import (
	"fmt"
	"bytes"
	"time"
	"math/rand"
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

func (o Order) String() string {
	var orderTypeString string

	if o.OrderType == BuyOrderType {
		orderTypeString = "Buy"
	} else {
		orderTypeString = "Sell"
	}

	return fmt.Sprintf("%v: %v shares of %v at $%v", orderTypeString, o.Shares, o. Symbol, o.Value)
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

// The symbols in the market
func (m Market) Symbols() []string {
	keys := make([]string, 0)
	for k, _ := range m.Stocks {
		keys = append(keys, k)
	}
	return keys	
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

func generateOrders(n int, m Market, orderChannel chan Order) {
	// Randomly sleep between orders
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		// Create a new order for a random stock with a random price
		symbols := m.Symbols()
		orderChannel <- Order{symbols[rand.Intn(len(symbols))], BuyOrderType, 100, 64.5}

		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
	close(orderChannel)
}

func startTrading(m *Market) {
	marketActivity := make(chan Order)

	go generateOrders(10, *m, marketActivity)
	for ord := range marketActivity {
		fmt.Printf("Got an order: %v\n", ord)
	}
}

func main() {
	market, _ := initializeMarket()
	fmt.Println(market)

	startTrading(&market)
}
