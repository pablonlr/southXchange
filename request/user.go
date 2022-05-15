package request

import (
	"encoding/json"
	"errors"

	"github.com/pablonlr/southXchange/types"
)

func (client *Client) NewDepositAddress(currency string) (string, error) {
	body := make(map[string]interface{})
	body["currency"] = currency
	resp, err := client.PostReq("generatenewaddress", body)
	if err != nil {
		return "", err
	}
	addr := ""
	err = json.Unmarshal(resp, &addr)
	if err != nil {
		return "", err
	}
	return addr, nil
}

func (client *Client) LNInvoice(currency string, amount float64) (string, error) {
	if currency != "BTC" && currency != "LTC" {
		return "", errors.New("Invalid currency")
	}
	body := make(map[string]interface{})
	body["currency"] = currency
	body["amount"] = amount
	resp, err := client.PostReq("getlninvoice", body)
	if err != nil {
		return "", err
	}
	invoice := ""
	err = json.Unmarshal(resp, &invoice)
	if err != nil {
		return "", err
	}
	return invoice, nil

}

func (client *Client) Withdraw(currency, address string, amount float64) (*types.Whithdrawal, error) {
	body := make(map[string]interface{})
	body["currency"] = currency
	body["address"] = address
	body["amount"] = amount
	resp, err := client.PostReq("withdraw", body)
	if err != nil {
		return nil, err
	}
	with := &types.Whithdrawal{}
	err = json.Unmarshal(resp, with)
	if err != nil {
		return nil, err
	}
	return with, nil

}

func (client *Client) Balances() ([]types.Balance, error) {
	body := make(map[string]interface{})
	resp, err := client.PostReq("listBalances", body)
	if err != nil {
		return nil, err
	}
	bls := []types.Balance{}
	err = json.Unmarshal(resp, &bls)
	if err != nil {
		return nil, err
	}
	return bls, nil
}

func (client *Client) Balance(currencies ...string) (map[string]float64, error) {
	mp := make(map[string]float64)
	bals, err := client.Balances()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(currencies); i++ {
		for j := 0; j < len(bals); j++ {
			if currencies[i] == bals[j].Currency {
				mp[currencies[i]] = bals[j].Available
				break
			}

		}
	}
	return mp, nil
}

func (client *Client) MakeOrder(currency, vscurrency, orderType string, amount, limitPrice float64) (string, error) {
	currency, vscurrency = cToUpper(currency, vscurrency)
	if orderType != "buy" && orderType != "sell" {
		return "", errors.New("Invaid orderType")
	}
	body := make(map[string]interface{})
	body["listingCurrency"] = currency
	body["referenceCurrency"] = vscurrency
	body["type"] = orderType
	body["amount"] = amount
	body["limitPrice"] = limitPrice
	resp, err := client.PostReq("placeOrder", body)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (client *Client) ListOrders() ([]types.PendingOrder, error) {
	body := make(map[string]interface{})
	resp, err := client.PostReq("listOrders", body)
	if err != nil {
		return nil, err
	}
	orders := []types.PendingOrder{}
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (client *Client) CancelOrder(orderCode string) error {
	body := make(map[string]interface{})
	body["orderCode"] = orderCode
	_, err := client.PostReq("cancelOrder", body)
	if err != nil {
		return err
	}
	return nil

}

func (client *Client) CancelAllOrders(currency, vscurrency string) error {
	body := make(map[string]interface{})
	body["listingCurrency"] = currency
	body["referenceCurrency"] = vscurrency
	_, err := client.PostReq("cancelMarketOrders", body)
	if err != nil {
		return err
	}
	return nil

}
func (client *Client) ListTransactions(currency, txtype string, page, pageSize int) ([]types.Transaction, error) {
	body := make(map[string]interface{})
	body["Currency"] = currency
	body["TransactionType"] = txtype
	//txtypes: deposits, withdrawal, transactions, depositwithdrawals
	body["PageIndex"] = page
	body["PageSize"] = pageSize
	resp, err := client.PostReq("listTransactions", body)
	if err != nil {
		return nil, err
	}
	txs := types.ListTransactions{}
	err = json.Unmarshal(resp, &txs)
	if err != nil {
		return nil, err
	}
	return txs.Result, nil
}
