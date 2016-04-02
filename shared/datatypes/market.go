package dt

import (
	"bytes"
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
