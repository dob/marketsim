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
import "github.com/dob/marketsim/shared/datatypes"

var market *dt.Market = dt.NewMarket()

// Add a stock or two.
// Lets add a symbols called AMZN and TSLA starting with the default price
market.stocks[dt.StockSymbol("AMZN")] = &dt.Stock{dt.StockSymbol("AMZN"), "Amazon", dt.StartingPrice}
market.stocks[dt.StockSymbol("TSLA")] = &dt.Stock{dt.StockSymbol("TSLA"), "Tesla", dt.StartingPrice}
```

Marketsim currently supports limit orders. Lets submit a few 100 share limit 
orders for AMZN.

``` go
buyOrder := dt.Order{"AMZN", dt.BuyOrderType, dt.LimitOrderType, 100, 645.20, dt.OrderStatusOpen}
sellOrder := dt.Order{"AMZN", dt.SellOrderType, dt.LimitOrderType, 100, 646.10, dt.OrderStatusOpen}

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
crossSpreadOrder := dt.Order{"AMZN", dt.SellOrderTYpe, dt.LimitOrderType, 50, 645.20, dt.OrderStatusOpen}
/// This order will cross the spread to the first submited sell order, and 50 shares will be taken off the OrderBook
```

## To-Do

1. Restructure the repo so that it follows standard library-first convention
2. Move the executable into an examples/cmd directory as it's not really anything other than a demonstration of usages
3. Support market orders and more sensible defaults.
4. Better error reporting
