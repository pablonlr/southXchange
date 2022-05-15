package types

import (
	"fmt"
	"strconv"
	"time"
)

type Market struct {
	Currency   string
	VsCurrency string
	ID         int
}

type Ticket struct {
	Bid      float64 `json:"Bid"`
	Ask      float64 `json:"Ask"`
	Last     float64 `json:"Last"`
	Var24    float64 `json:"Variation24Hr"`
	Volume24 float64 `json:"Volume24Hr"`
}

type Order struct {
	Index  int     `json:"Index"`
	Amount float64 `json:"Amount"`
	Price  float64 `json:"Price"`
}

type OrderBook struct {
	BuyOrders  []Order `json:"BuyOrders"`
	SellOrders []Order `json:"SellOrders"`
}

func (book *OrderBook) String(count, decimalsQ, decimalsR int) string {
	result := "\nSells:\n"
	result += addst(reverse(book.SellOrders[:min(len(book.SellOrders), count)]), decimalsQ, decimalsR)
	result += "\nBuys:\n"
	result += addst(book.BuyOrders[:min(len(book.BuyOrders), count)], decimalsQ, decimalsR)
	return result
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func reverse(ords []Order) []Order {
	od := make([]Order, len(ords))
	for i, x := range ords {
		od[len(od)-i-1] = x
	}
	return od
}
func addst(orders []Order, decimalsQ, decimalsR int) (result string) {
	for i := 0; i < len(orders); i++ {
		quantSt := strconv.FormatFloat(orders[i].Amount, 'f', decimalsQ, 64)
		rate := strconv.FormatFloat(orders[i].Price, 'f', decimalsR, 64)
		result += fmt.Sprintf("%s => %s\n", quantSt, rate)
	}
	return
}

type PastTrades struct {
	At     int64   `json:"At"`
	Amount float64 `json:"Amount"`
	Price  float64 `json:"Price"`
	Type   string  `json:"Type"`
}

type Balance struct {
	Currency    string  `json:"Currency"`
	Deposited   float64 `json:"Deposited"`
	Available   float64 `json:"Available"`
	Unconfirmed float64 `json:"Unconfirmed"`
}

type PendingOrder struct {
	Code              string  `json:"Code"`
	Type              string  `json:"Type"`
	Amount            float64 `json:"Amount"`
	OriginalAmount    float64 `json:"OriginalAmount"`
	LimitPrice        float64 `json:"LimitPrice"`
	ListingCurrency   string  `json:"ListingCurrency"`
	ReferenceCurrency string  `json:"ReferenceCurrency"`
}

type Whithdrawal struct {
	Status     string  `json:"Status"`
	Max        float64 `json:"Max"`
	MaxDaily   float64 `json:"MaxDaily"`
	MovementId int64   `json:"MovementId"`
}

const layoutISO = "2006-01-02T15:04:05"

type Transaction struct {
	Date          string  `json:"Date"`
	CurrencyCode  string  `json:"CurrencyCode"`
	Amount        float64 `json:"Amount"`
	TotalBalance  float64 `json:"TotalBalance"`
	Type          string  `json:"Type"`
	Status        string  `json:"Status"`
	Address       string  `json:"Address"`
	Hash          string  `json:"Hash"`
	Price         float64 `json:"Price"`
	OtherAmount   float64 `json:"OtherAmount"`
	OtherCurrency string  `json:"OtherCurrency"`
	OrderCode     string  `json:"OrderCode"`
	TradeId       uint64  `json:"TradeId"`
	MovementId    uint64  `json:"MovementId"`
}

func (tx *Transaction) GetDate() (time.Time, error) {
	t, err := time.Parse(layoutISO, tx.Date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

type ListTransactions struct {
	TotalElements int           `json:"TotalElements"`
	Result        []Transaction `json:"Result"`
}

func FilterTXs(txs []Transaction, txType string, afterDate time.Time) ([]Transaction, error) {
	arr := []Transaction{}
	for _, v := range txs {
		t, err := v.GetDate()
		if err != nil {
			return nil, err
		}
		if v.Type == txType && t.After(afterDate) {
			arr = append(arr, v)
		}
	}
	return arr, nil
}
