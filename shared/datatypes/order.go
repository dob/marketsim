package dt

import (
	"fmt"
)

// Represents an order to buy or sell
type Order struct {
	Symbol string
	BuySell OrderBuySellVal
	OrderType OrderTypeVal
	Shares int
	Value float64
	OrderStatus
}

type OrderBuySellVal int
const (
	BuyOrderType OrderBuySellVal = iota + 1
	SellOrderType
)

type OrderTypeVal int
const (
	MarketOrderType OrderTypeVal = iota + 1
	LimitOrderType
)

type OrderStatus int
const (
	OrderStatusOpen OrderStatus = iota + 1
	OrderStatusFilled
	OrderStatusPartial
	OrderStatusCancelled
)

func (o OrderTypeVal) String() string {
	switch o {
	case MarketOrderType:
		return "Market"
	case LimitOrderType:
		return "Limit"
	default:
		return "Unknown"
	}
}


func (o Order) String() string {
	var buySellTypeString string

	if o.BuySell == BuyOrderType {
		buySellTypeString = "Buy"
	} else {
		buySellTypeString = "Sell"
	}

	return fmt.Sprintf("%v: %v shares of %v at $%v. %v order.", buySellTypeString, o.Shares, o. Symbol, o.Value, o.OrderType)
}

type SortedOrders []*Order
// Sort interface functions
func (s SortedOrders) Len() int {
	return len(s)
}

func (s SortedOrders) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortedOrders) Less(i, j int) bool {
	return s[i].Value < s[j].Value
}
