package request

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/pablonlr/southXchange/types"
)

func (client *Client) Markets() ([]types.Market, error) {
	resp, err := client.GetReq("v2", "markets")
	if err != nil {
		return nil, err
	}
	var arr []([]interface{})
	err = json.Unmarshal(resp, &arr)
	if err != nil {
		return nil, err
	}
	marketArr := make([]types.Market, len(arr))
	for i, v := range arr {
		curr, vscurr, id, ok := getMarket(v)
		if !ok {
			return nil, errors.New("Error durning type assertion")
		}
		mk := types.Market{
			Currency:   curr,
			VsCurrency: vscurr,
			ID:         id,
		}
		marketArr[i] = mk
	}
	return marketArr, nil

}
func getMarket(arr []interface{}) (curr string, vscurr string, id int, ok bool) {
	ok = true
	curr, ok = arr[0].(string)
	if !ok {
		return "", "", -1, false
	}
	vscurr, ok = arr[1].(string)
	if !ok {
		return "", "", -1, false
	}
	f := arr[2].(float64)
	if !ok {
		return "", "", -1, false
	}
	id = int(f)
	return
}

func (client *Client) Price(coin, vscoin string) (*types.Ticket, error) {
	coin, vscoin = cToUpper(coin, vscoin)
	resp, err := client.GetReq("price", coin, vscoin)
	if err != nil {
		return nil, err
	}
	ticker := &types.Ticket{}
	err = json.Unmarshal(resp, ticker)
	if err != nil {
		return nil, err
	}
	return ticker, nil
}
func cToUpper(coin, vscoin string) (string, string) {
	coin = strings.ToUpper(coin)
	vscoin = strings.ToUpper(vscoin)
	return coin, vscoin
}
func (client *Client) OrderBook(coin, vscoin string) (*types.OrderBook, error) {
	coin, vscoin = cToUpper(coin, vscoin)
	resp, err := client.GetReq("book", coin, vscoin)
	if err != nil {
		return nil, err
	}
	book := &types.OrderBook{}
	err = json.Unmarshal(resp, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
func (client *Client) PastTrades(coin, vscoin string) ([]types.PastTrades, error) {
	coin, vscoin = cToUpper(coin, vscoin)
	resp, err := client.GetReq("trades", coin, vscoin)
	if err != nil {
		return nil, err
	}
	trades := []types.PastTrades{}
	err = json.Unmarshal(resp, &trades)
	if err != nil {
		return nil, err
	}
	return trades, nil

}
