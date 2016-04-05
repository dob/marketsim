package main

import (
	"fmt"
	"time"
	"math/rand"

	"github.com/dob/marketsim/shared/datatypes"
	"github.com/dob/marketsim/shared/utils"
)

// START SIMULATION

// Populate the market with some fake data
func stubMarketStocks(m *dt.Market) {
	//m.Stocks = make(map[dt.StockSymbol]*dt.Stock)

	sym1 := dt.Stock{"AMZN", "Amazon", dt.StartingPrice}
	m.Stocks[sym1.Symbol] = &sym1

	sym2 := dt.Stock{"TSLA", "Tesla Motors", dt.StartingPrice}
	m.Stocks[sym2.Symbol] = &sym2

	sym3 := dt.Stock{"TWTR", "Twitter", dt.StartingPrice}
	m.Stocks[sym3.Symbol] = &sym3
}

// Initialize the market
func initializeMarketWithStocks() (*dt.Market, error) {
	var market *dt.Market = dt.NewMarket()
	stubMarketStocks(market)

	return market, nil
}

//Randomly generate n orders and sleep for 0-2 seconds in between
func generateOrders(n int, m *dt.Market, orderChannel chan *dt.Order) {
	// Randomly sleep between orders
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		// Create a new order for a random stock with a random price
		symbols := m.Symbols()

		// Generate the order params
		symbol := symbols[rand.Intn(len(symbols))]
		buySellType := dt.OrderBuySellVal(rand.Intn(2) + 1) // Will be either buy or sell (1 or 2)
		orderType := dt.LimitOrderType //dt.OrderTypeVal(rand.Intn(2) + 1) // Will be either market or limit (1 or 2)

		var price float64
		priceForSymbol := m.Stocks[symbol].Price
		if orderType == dt.LimitOrderType {
			if ((priceForSymbol.Bid == dt.StartingPrice.Bid) || (priceForSymbol.Offer == dt.StartingPrice.Offer)) {
				// Go pretty random unless we have both bid and ask
				price = float64(rand.Intn(100)) + rand.Float64()
			} else {
				// Randomly choose positive or negative price fluctuation
				// within 5% of the current price
				priceDelta := float64(priceForSymbol.MidPrice() * 0.05) * rand.Float64()
				
				if rand.Intn(2) == 0 {
					price = priceForSymbol.Bid + priceDelta
				} else {
					price = priceForSymbol.Bid - priceDelta
				}
			}
		}

		price = utils.RoundToPlaces(price, 2)
		shares := (rand.Intn(10) + 1) * 100  // Start with 100x share lots

		order := dt.Order{symbol, buySellType, orderType, shares, price, dt.OrderStatusOpen}			
		
		orderChannel <- &order

		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	}
	close(orderChannel)
}

// Simulate the start of trading
func startTrading(m *dt.Market) {
	marketActivity := make(chan *dt.Order)

	go generateOrders(1000, m, marketActivity)
	for ord := range marketActivity {
		//log.Printf("Got an order: %v\n", ord)
		m.ReceiveOrder(ord)
	}
}

func main() {
	market, _ := initializeMarketWithStocks()
	fmt.Println(market)
	startTrading(market)
	fmt.Println(market)
}
