# Marketsim

Marketsim is a library to aid in simulation of a market. It allows you to create
a market, add stocks to your market, submit limit orders to the market, 
and then keep track of the execution of trades that occur if they can be
facilitated by the prices and quantities set in the limit orders. The bid/ask 
spread for the stock will be updated according to the outstanding orders
on the order book.

*(This is a work in progress and isn't ready for anything other than experimenation yet).*

## Installation

The easiest way to install marketsim is to run `go get github.com/dob/marketsim`. 

## Usage

To use marketsim, you first have to instantiate a market and add at least one stock to it.

``` go
import ms "github.com/dob/marketsim"

var market *ms.Market = ms.NewMarket()

// Add a stock or two.
// Lets add a symbols called AMZN and TSLA starting with the default price
market.Stocks[ms.StockSymbol("AMZN")] = &ms.Stock{ms.StockSymbol("AMZN"), "Amazon", ms.StartingPrice}
market.Stocks[ms.StockSymbol("TSLA")] = &ms.Stock{ms.StockSymbol("TSLA"), "Tesla", ms.StartingPrice}
```

Marketsim currently supports limit orders. Lets submit a few 100 share limit 
orders for AMZN.

``` go
buyOrder := &ms.Order{"AMZN", ms.BuyOrderType, ms.LimitOrderType, 100, 645.20, ms.OrderStatusOpen}
sellOrder := &ms.Order{"AMZN", ms.SellOrderType, ms.LimitOrderType, 100, 646.10, ms.OrderStatusOpen}

market.ReceiveOrder(buyOrder)
market.ReceiveOrder(sellOrder)
```

Now we should have two orders on the OrderBook. You can check the price of 
the orders with 

``` go
price := market.GetPriceForSymbol("AMZN")

fmt.Printf("Bid price: %v, Ask price: %v", price.Bid, price.Offer)
```

You can also print the market object directly which will print a nice tabled 
output of all prices on the market. `fmt.println(market)`

If an order comes in that crosses the spread, the market will process the order,
update the order book, and prices accordingly.

``` go
crossSpreadOrder := ms.Order{"AMZN", ms.SellOrderType,
ms.LimitOrderType, 50, 645.20, ms.OrderStatusOpen}
market.ReceiveOrder(crossSpreadOrder)

// This order will cross the spread to the first submited sell order, and 50 shares will be taken off the OrderBook
```

## To-Do

1. Move the executable into an examples/cmd directory as it's not really anything other than a demonstration of usages
1. Support market orders and more sensible defaults.
1. Better error reporting
