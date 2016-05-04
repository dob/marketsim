package dt

// Represent the current price of a stock
type StockPrice struct {
	Bid   float64
	Offer float64
}

func (sp StockPrice) MidPrice() float64 {
	if (sp.Bid == StartingPrice.Bid) && (sp.Offer == StartingPrice.Offer) {
		return 0
	} else if sp.Bid == MinPrice {
		return MaxPrice
	} else if sp.Offer == MaxPrice {
		return MinPrice
	} else {
		return sp.Bid + ((sp.Offer - sp.Bid) / 2.0)
	}
}

const MinPrice float64 = 0.0
const MaxPrice float64 = float64(1 << 10) // Max price in our market is $1M+/share

var StartingPrice StockPrice = StockPrice{MinPrice, MaxPrice}
