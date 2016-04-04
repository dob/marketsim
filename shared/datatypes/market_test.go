package dt

import (
	"testing"
	"os"
)

var m Market

func TestMain(m *testing.M) {
	setup()

	retval := m.Run()

	os.Exit(retval)
}

func setup() {
	m = Market{}
	m.Stocks = make(map[StockSymbol]*Stock)
	m.Orders = make(map[StockSymbol][]*Order)
	m.Stocks["AMZN"] = &Stock{"AMZN", "Amazon", StockPrice{}}
	m.Stocks["TSLA"] = &Stock{"TSLA", "Tesla", StockPrice{}}
	m.Stocks["TWTR"] = &Stock{"TWTR", "Twitter", StockPrice{}}
}

func TestReceiveOrderNegativeShares(t *testing.T) {
	o := Order{"AMZN", BuyOrderType, MarketOrderType, -2, 23.0, OrderStatusOpen}
	err := m.ReceiveOrder(o)

	if err == nil {
		t.Errorf("ReceiveOrder didn't fail on negative shares")
	}

	o1 := Order{"TSLA", BuyOrderType, MarketOrderType, 23, 23.0, OrderStatusOpen}
	err = m.ReceiveOrder(o1)

	if err != nil {
		t.Errorf("ReceiveOrder failed on a valid order")
	}
}

func TestReceiveMarketOrders(t *testing.T) {
	t.Skip()
}

func TestSetPrices(t *testing.T) {
	t.Skip()
}
