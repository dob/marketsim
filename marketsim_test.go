package marketsim

import (
	"os"
	"testing"
)

var m *Market

func TestMain(testingm *testing.M) {
	setup()

	retval := testingm.Run()

	os.Exit(retval)
}

func setup() {
	m = &Market{}
	m.Stocks = make(map[StockSymbol]*Stock)
	m.Orders = make(map[StockSymbol][]*Order)
	m.Stocks["AMZN"] = &Stock{"AMZN", "Amazon", StockPrice{}}
	m.Stocks["TSLA"] = &Stock{"TSLA", "Tesla", StockPrice{}}
	m.Stocks["TWTR"] = &Stock{"TWTR", "Twitter", StockPrice{}}
}

func TestReceiveOrderNegativeShares(t *testing.T) {
	o := Order{"AMZN", BuyOrderType, MarketOrderType, -2, 23.0, OrderStatusOpen}
	err := m.ReceiveOrder(&o)

	if err == nil {
		t.Errorf("ReceiveOrder didn't fail on negative shares")
	}

	o1 := Order{"TSLA", BuyOrderType, MarketOrderType, 23, 23.0, OrderStatusOpen}
	err = m.ReceiveOrder(&o1)

	if err != nil {
		t.Errorf("ReceiveOrder failed on a valid order")
	}
}

func TestReceiveOrdersDifferentStocks(t *testing.T) {
	setup()
	order := Order{"AMZN", BuyOrderType, LimitOrderType, 100, 595.0, OrderStatusOpen}
	order2 := Order{"TSLA", SellOrderType, LimitOrderType, 100, 596.50, OrderStatusOpen}
	m.ReceiveOrder(&order)
	m.ReceiveOrder(&order2)

	if len(m.Orders) != 2 {
		t.Errorf("Orders were clearly not received as length was %v", len(m.Orders))
	}
}

func TestReceiveOrdersSameStocks(t *testing.T) {
	setup()
	order := Order{"AMZN", BuyOrderType, LimitOrderType, 100, 595.0, OrderStatusOpen}
	order2 := Order{"AMZN", SellOrderType, LimitOrderType, 100, 596.50, OrderStatusOpen}
	m.ReceiveOrder(&order)
	m.ReceiveOrder(&order2)

	if len(m.Orders["AMZN"]) != 2 {
		t.Errorf("Orders were clearly not received as length was %v", len(m.Orders))
	}

}

func TestSetPrices(t *testing.T) {
	setup()
	order := Order{"AMZN", BuyOrderType, LimitOrderType, 100, 595.0, OrderStatusOpen}
	order2 := Order{"AMZN", SellOrderType, LimitOrderType, 100, 596.50, OrderStatusOpen}
	m.ReceiveOrder(&order)
	m.ReceiveOrder(&order2)

	AMZNPrice := m.Stocks["AMZN"].Price
	if AMZNPrice.Bid != 595.0 {
		t.Errorf("Bid price was supposed to be 595 but instead was %v", AMZNPrice.Bid)
	}

	if AMZNPrice.Offer != 596.50 {
		t.Errorf("Offer price was supposed to be 596.5 but instead was %v", AMZNPrice.Bid)
	}
}

func TestIfOrderWillProcess(t *testing.T) {
	setup()
	order := Order{"AMZN", BuyOrderType, LimitOrderType, 100, 595.0, OrderStatusOpen}
	order2 := Order{"AMZN", SellOrderType, LimitOrderType, 100, 595.0, OrderStatusOpen}
	m.ReceiveOrder(&order)
	m.ReceiveOrder(&order2)

	amznOrders := m.Orders["AMZN"]
	if len(amznOrders) != 0 {
		t.Errorf("Orders should have been processed and removed but there are %v", len(amznOrders))
	}

	if order.OrderStatus != OrderStatusFilled {
		t.Errorf("Order should have been filled but was not")
	}

	if order2.OrderStatus != OrderStatusFilled {
		t.Errorf("Order should have been filled but was not")
	}
}
