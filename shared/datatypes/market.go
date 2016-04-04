package dt

import (
	"bytes"
	"errors"
	"math"
	"sort"
)

// Establish the market
type Market struct {
	Stocks map[string]*Stock
	Orders []*Order
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
func (m *Market) Symbols() []string {
	keys := make([]string, 0)
	for k, _ := range m.Stocks {
		keys = append(keys, k)
	}
	return keys	
}

var InvalidOrder = errors.New("Invalid order")

func (m *Market) ReceiveOrder(o Order) error {
	// Check for errors
	if o.Shares < 1 {
		return InvalidOrder
	}

	// Put the order into the market
	m.Orders = append(m.Orders, &o)

	// Process the order
	m.processOrder(&o)
	
	// Update the price of the stock represented by the order o
	m.updatePriceForOrder(o)
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
	
}

func (m *Market) processMarketOrder(o *Order) {
	sharesOutstanding := o.Shares
	potentialOrders := m.getOrdersOnOtherSide(o)

	// Find the highest or lowest prices and fulfill the orders
	for _, candidateOrder := range potentialOrders {
		if sharesOutstanding <= candidateOrder.Shares {
			// fullfill the full order and part of the candidate order
			m.fullfillOrder(o, sharesOutstanding, candidateOrder.Value)
			m.fullfillOrder(candidateOrder, candidateOrder.Shares - sharesOutstanding, candidateOrder.Value)
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
	potentialOrders := make([]*Order, 0)

	for _, potentialOrder := range m.Orders {
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
	for i, ords := range m.Orders {
		if ords == o {
			m.Orders = append(m.Orders[:i], m.Orders[i+1:]...)
			break
		}
	}
}

func (m *Market) updatePriceForOrder(o Order) {
	// Find the appropriate stock
	s := m.Stocks[o.Symbol]
	currentBid := s.Price.Bid
	currentOffer := s.Price.Offer

	if o.BuySell == BuyOrderType {
		if currentBid == 0 {
			s.Price.Bid = math.Max(currentBid, currentOffer)
		}

		if o.Value > currentBid {
			s.Price.Bid = math.Max(o.Value, currentOffer)
		}
		
	} else {

	}
}
