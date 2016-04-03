package dt

import (
	"fmt"
)

// Equity representing one stock
type Stock struct {
	Symbol string
	Name string
	Price StockPrice
}

func (s Stock) String() string {
	return fmt.Sprintf("%v, %v: %v-%v", s.Symbol, s.Name, s.Price.Bid, s.Price.Offer)
}


