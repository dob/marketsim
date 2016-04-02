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
}

type OrderBuySellVal int
const (
	BuyOrderType OrderBuySellVal = iota
	SellOrderType OrderBuySellVal = iota
)

type OrderTypeVal int
const (
	MarketOrderType OrderTypeVal = iota
	LimitOrderType OrderTypeVal = iota
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
