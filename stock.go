package marketsim

import (
	"fmt"
	"github.com/leekchan/accounting"
)

type StockSymbol string

// Equity representing one stock
type Stock struct {
	Symbol StockSymbol
	Name   string
	Price  StockPrice
}

func (s Stock) String() string {
	return fmt.Sprintf("%v, %v: %v-%v", s.Symbol, s.Name, cf.FormatMoney(s.Price.Bid), cf.FormatMoney(s.Price.Offer))
}

var cf accounting.Accounting = accounting.Accounting{Symbol: "$", Precision: 2}
