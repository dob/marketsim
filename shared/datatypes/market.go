package dt

import (
	"bytes"
	"errors"
	"math"
	"sort"
	"sync"

	"github.com/dob/marketsim/shared/utils"
)

// Establish the market
type Market struct {
	Stocks map[StockSymbol]*Stock
	Orders map[StockSymbol][]*Order

	mux sync.Mutex   // Make the market threadsafe
}

func NewMarket() *Market {
	m := &Market{}
	m.Stocks = make(map[StockSymbol]*Stock)
	m.Orders = make(map[StockSymbol][]*Order)
	return m
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
func (m *Market) Symbols() []StockSymbol {
	keys := make([]StockSymbol, 0)
	for k, _ := range m.Stocks {
		keys = append(keys, k)
	}
	return keys	
}

func (m *Market) GetPriceForSymbol(s StockSymbol) StockPrice {
	m.mux.Lock()
	price := m.Stocks[s].Price
	m.mux.Unlock()
	
	return price
}

var InvalidOrder = errors.New("Invalid order")

func (m *Market) ReceiveOrder(o *Order) error {
	// Make sure that when we're updating the market nothing else
	// can read from the market
	m.mux.Lock()
	defer m.mux.Unlock()
	
	// Check for errors
	if o.Shares < 1 {
		return InvalidOrder
	}

	// Put the order into the market
	m.Orders[o.Symbol] = append(m.Orders[o.Symbol], o)
	//log.Printf("Adding order: %v\n", o)

	// Process the order
	m.processOrder(o)
	
	// Update the price of the stock represented by the order o
	m.updatePriceForSymbol(o.Symbol)
	return nil
}

func (m *Market) processOrder(o *Order) {
	switch o.OrderType {
	case MarketOrderType:
		m.processMarketOrder(o)
	case LimitOrderType:
		m.processLimitOrder(o)
	}
}

func (m *Market) processLimitOrder(o *Order) {
	sharesOutstanding := o.Shares
	potentialOrders := m.getOrdersOnOtherSide(o)

	// Find the highest or lowest prices and fulfill the orders
	for _, candidateOrder := range potentialOrders {
		if (o.BuySell == BuyOrderType && candidateOrder.Value > o.Value) ||
			(o.BuySell == SellOrderType && candidateOrder.Value < o.Value) {
			// The offer doesn't match the limit criteria so abort and
			// leave this offer around
			break
		}
		
		if sharesOutstanding <= candidateOrder.Shares {
			// fullfill the full order and part of the candidate order
			m.fullfillOrder(o, sharesOutstanding, candidateOrder.Value)
			m.fullfillOrder(candidateOrder, sharesOutstanding, candidateOrder.Value)
			sharesOutstanding = 0
			break
		} else {
			// partially fullfill the order and continue to the next order
			m.fullfillOrder(o, candidateOrder.Shares, candidateOrder.Value)
			m.fullfillOrder(candidateOrder, candidateOrder.Shares, candidateOrder.Value)
			sharesOutstanding -= candidateOrder.Shares
		}
	}	
}

func (m *Market) processMarketOrder(o *Order) {
	sharesOutstanding := o.Shares
	potentialOrders := m.getOrdersOnOtherSide(o)

	// Find the highest or lowest prices and fulfill the orders
	for _, candidateOrder := range potentialOrders {
		if sharesOutstanding <= candidateOrder.Shares {
			// fullfill the full order and part of the candidate order
			m.fullfillOrder(o, sharesOutstanding, candidateOrder.Value)
			m.fullfillOrder(candidateOrder, sharesOutstanding, candidateOrder.Value)
			sharesOutstanding = 0
			break
		} else {
			// partially fullfill the order and continue to the next order
			m.fullfillOrder(o, candidateOrder.Shares, candidateOrder.Value)
			m.fullfillOrder(candidateOrder, candidateOrder.Shares, candidateOrder.Value)
			sharesOutstanding -= candidateOrder.Shares
		}
	}
}

func (m *Market) getOrdersOnOtherSide(o *Order) []*Order {
	// Make a copy of the potential orders so we can sort
	potentialOrders := make([]*Order, 0)

	for _, potentialOrder := range m.Orders[o.Symbol] {
		if o.Symbol == potentialOrder.Symbol &&
			o.BuySell != potentialOrder.BuySell {
			potentialOrders = append(potentialOrders, potentialOrder)
		}
	}

	// Sort the orders by price. Reverse if you're looking to sell
	if o.BuySell == SellOrderType {
		sort.Sort(sort.Reverse(SortedOrders(potentialOrders)))
	} else {
		sort.Sort(SortedOrders(potentialOrders))
	}

	return potentialOrders
}

func (m *Market) fullfillOrder(o *Order, shares int, price float64) {
	// Should log what happened here

	if o.Shares == shares {
		o.OrderStatus = OrderStatusFilled
		m.removeOrder(o)
	} else {
		o.Shares = o.Shares - shares
		o.OrderStatus = OrderStatusPartial
	}
}

func (m *Market) removeOrder(o *Order) {
	ordersForSymbol := m.Orders[o.Symbol]
	for i, ords := range ordersForSymbol {
		if ords == o {
			ordersForSymbol = append(ordersForSymbol[:i], ordersForSymbol[i+1:]...)
			m.Orders[o.Symbol] = ordersForSymbol
			//log.Printf("Removing order %v and now there are %v", o, len(ordersForSymbol))

			break
		}
	}
}

func (m *Market) updatePriceForSymbol(ss StockSymbol) {
	// Find the appropriate stock
	stock := m.Stocks[ss]
	orders := m.Orders[ss]

	maxBid := MinPrice
	minOffer := MaxPrice

	for _, o := range orders {
		// Only update for limit orders
		if o.OrderType == LimitOrderType {
			if o.BuySell == BuyOrderType {
				maxBid = math.Max(o.Value, maxBid)
			} else {
				minOffer = math.Min(o.Value, minOffer)
			}
		}
	}

	stock.Price.Bid = utils.RoundToPlaces(maxBid, 2)
	stock.Price.Offer = utils.RoundToPlaces(minOffer, 2)

	//log.Printf("Updating the price of %v to $%v - $%v", ss, maxBid, minOffer)
	// Write the stock back into the market? why do you have to do this
	// since stock should already be a pointer to the stock struct
	m.Stocks[ss] = stock
}

