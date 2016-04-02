package main

import (
	"fmt"
	"time"
	"math/rand"

	"github.com/dob/marketsim/shared/datatypes"
)

// START SIMULATION

// Populate the market with some fake data
func stubMarketStocks(m *dt.Market) {
	m.Stocks = make(map[string]dt.Stock)

	sym1 := dt.Stock{"AMZN", "Amazon", 568.2, 568.4}
	m.Stocks[sym1.Symbol] = sym1

	sym2 := dt.Stock{"TSLA", "Tesla Motors", 288.5, 289}
	m.Stocks[sym2.Symbol] = sym2

	sym3 := dt.Stock{"TWTR", "Twitter", 15.4, 15.60}
	m.Stocks[sym3.Symbol] = sym3
}

// Initialize the market
func initializeMarket() (dt.Market, error) {
	var market dt.Market
	stubMarketStocks(&market)
	market.Orders = make([]dt.Order, 0)

	return market, nil
}

//Randomly generate n orders and sleep for 0-2 seconds in between
func generateOrders(n int, m dt.Market, orderChannel chan dt.Order) {
	// Randomly sleep between orders
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		// Create a new order for a random stock with a random price
		symbols := m.Symbols()
		orderChannel <- dt.Order{symbols[rand.Intn(len(symbols))], dt.BuyOrderType, dt.LimitOrderType, 100, 64.5}

		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
	close(orderChannel)
}

// Simulate the start of trading
func startTrading(m *dt.Market) {
	marketActivity := make(chan dt.Order)

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
