package main

import (
	"fmt"
	ms "github.com/dob/marketsim"
)

type SimpleExample struct{}

func (s SimpleExample) run() {
	var market *ms.Market = ms.NewMarket()

	// Add a stock or two.
	// Lets add a symbols called AMZN and TSLA starting with the default price
	market.Stocks[ms.StockSymbol("AMZN")] = &ms.Stock{Symbol: ms.StockSymbol("AMZN"), Name: "Amazon", Price: ms.StartingPrice}
	market.Stocks[ms.StockSymbol("TSLA")] = &ms.Stock{Symbol: ms.StockSymbol("TSLA"), Name: "Tesla", Price: ms.StartingPrice}

	// Sumit two orders into the market
	buyOrder := &ms.Order{"AMZN", ms.BuyOrderType, ms.LimitOrderType, 100, 645.20, ms.OrderStatusOpen}
	sellOrder := &ms.Order{"AMZN", ms.SellOrderType, ms.LimitOrderType, 100, 646.10, ms.OrderStatusOpen}

	market.ReceiveOrder(buyOrder)
	market.ReceiveOrder(sellOrder)

	// List Prices
	price := market.GetPriceForSymbol("AMZN")

	fmt.Printf("Bid price: %v, Ask price: %v\n", price.Bid, price.Offer)

	// This order will cross the spread to the first submited sell order, and 50 shares will be taken off the OrderBook
	crossSpreadOrder := &ms.Order{"AMZN", ms.SellOrderType, ms.LimitOrderType, 50, 645.20, ms.OrderStatusOpen}
	market.ReceiveOrder(crossSpreadOrder)

	fmt.Println(market)
}
