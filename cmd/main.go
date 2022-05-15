package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pablonlr/southxchange/request"
)

const layoutISO = "2006-01-02T15:04:05"

func main() {
	client := &request.Client{http.Client{Timeout: 10 * time.Second}}
	//resp, err := client.MakeOrder("CRW", "BTC", "sell", 1, 0.000004)
	//resp, err := client.Balances()
	//resp, err := client.NewDepositAddress("CRW")
	//resp, err := client.ListOrders()
	//err := client.CancelOrder("95413567")
	resp, err := client.ListTransactions("CRW", "transactions", 0, 10)
	if err != nil {
		panic(err)
	}
	resp[0].GetDate()
	t, err := time.Parse(layoutISO, resp[0].Date)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println(t)

}
