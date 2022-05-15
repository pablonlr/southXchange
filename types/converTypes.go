package types

import "github.com/pablonlr/exchange"

func (book *OrderBook) ConvertOrderBook() *exchange.OrderBook {
	exbook := exchange.OrderBook{}
	exbook.Bid = fillOrders(book.BuyOrders, []exchange.Order{})
	exbook.Ask = fillOrders(book.SellOrders, []exchange.Order{})
	return &exbook
}

func fillOrders(from []Order, to []exchange.Order) []exchange.Order {
	for _, x := range from {
		newOrder := exchange.Order{Quantity: x.Amount, Price: x.Price}
		to = append(to, newOrder)
	}
	return to
}
