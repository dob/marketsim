package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	ms "github.com/dob/marketsim"
	"github.com/dob/marketsim/utils"
)

const NUMBER_OF_ORDERS_IN_SIMULATION = 100000
const STOCK_SEED_FILE_LOC = "data/seed/nasdaq_stocks.csv"

// START SIMULATION

// Populate the market with some fake data
func stubMarketStocks(m *ms.Market) {
	stocks := loadStocksFromSeedFile()
	for _, stock := range stocks {
		m.Stocks[stock.Symbol] = stock
	}
}

func loadStocksFromSeedFile() []*ms.Stock {
	stocks := make([]*ms.Stock, 0)

	// Find path for seed file relative to where we're at
	basePath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "dob", "marketsim")
	csvFile, _ := os.Open(filepath.Join(basePath, STOCK_SEED_FILE_LOC))
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	rawCSVData, _ := reader.ReadAll()

	for _, each := range rawCSVData {
		sym, name := each[0], each[1]
		fmt.Printf("adding %v: %v\n", sym, name)
		stocks = append(stocks, &ms.Stock{ms.StockSymbol(sym), name, ms.StartingPrice})
	}

	return stocks
}

// Initialize the market
func initializeMarketWithStocks() (*ms.Market, error) {
	var market *ms.Market = ms.NewMarket()
	stubMarketStocks(market)

	return market, nil
}

//Randomly generate n orders and sleep for 0-2 seconds in between
func generateOrders(n int, m *ms.Market, orderChannel chan *ms.Order) {
	// Randomly sleep between orders
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		// Create a new order for a random stock with a random price
		symbols := m.Symbols()

		// Generate the order params
		symbol := symbols[rand.Intn(len(symbols))]
		buySellType := ms.OrderBuySellVal(rand.Intn(2) + 1) // Will be either buy or sell (1 or 2)
		orderType := ms.LimitOrderType                      // Right now we're only seeding orders as limit types to set prices

		var price float64

		priceForSymbol := m.GetPriceForSymbol(symbol)

		if orderType == ms.LimitOrderType {
			if (priceForSymbol.Bid == ms.StartingPrice.Bid) || (priceForSymbol.Offer == ms.StartingPrice.Offer) {
				// Go pretty random unless we have both bid and ask
				price = float64(rand.Intn(100)) + rand.Float64()
			} else {
				// Randomly choose positive or negative price fluctuation
				// within 5% of the current price
				priceDelta := float64(priceForSymbol.MidPrice()*0.05) * rand.Float64()

				if rand.Intn(2) == 0 {
					price = priceForSymbol.Bid + priceDelta
				} else {
					price = priceForSymbol.Bid - priceDelta
				}
			}
		}

		price = utils.RoundToPlaces(price, 2)
		shares := (rand.Intn(10) + 1) * 100 // Start with 100x share lots

		order := ms.Order{symbol, buySellType, orderType, shares, price, ms.OrderStatusOpen}

		orderChannel <- &order

		// Do we want to sleep for an interval between orders?
		//time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	}
	close(orderChannel)
}

// Simulate the start of trading
func startTrading(m *ms.Market) {
	marketActivity := make(chan *ms.Order)

	go generateOrders(NUMBER_OF_ORDERS_IN_SIMULATION, m, marketActivity)
	for ord := range marketActivity {
		m.ReceiveOrder(ord)
	}
}

func main() {
	market, _ := initializeMarketWithStocks()
	fmt.Println(market)
	startTrading(market)
	fmt.Println(market)
}
