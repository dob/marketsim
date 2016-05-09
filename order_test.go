package marketsim

import (
	"sort"
	"testing"
)

func TestOrderSorting(t *testing.T) {
	orders := []*Order{&Order{"AMZN", BuyOrderType, MarketOrderType, 34, 23.0, OrderStatusOpen},
		&Order{"TSLA", BuyOrderType, MarketOrderType, 50, 203.0, OrderStatusOpen},
		&Order{"WLB", BuyOrderType, MarketOrderType, 60, 2.0, OrderStatusOpen}}

	sort.Sort(SortedOrders(orders))
	if orders[0].Symbol != "WLB" || orders[1].Symbol != "AMZN" || orders[2].Symbol != "TSLA" {
		t.Errorf("Sort did not work.")
	}

	sort.Sort(sort.Reverse(SortedOrders(orders)))

	if orders[0].Symbol != "TSLA" || orders[1].Symbol != "AMZN" || orders[2].Symbol != "WLB" {
		t.Errorf("Sort reverse did not work: %v", orders)
	}
}
