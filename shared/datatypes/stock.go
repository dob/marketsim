package dt

import (
	"fmt"
)

// Equity representing one stock
type Stock struct {
	Symbol string
	Name string
	Bid float64
	Offer float64
}

func (s Stock) String() string {
	return fmt.Sprintf("%v, %v: %v-%v", s.Symbol, s.Name, s.Bid, s.Offer)
}


