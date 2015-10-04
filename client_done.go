package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

type StockRequest struct {
	Budget                   float64 `json:"budget"`
	StockSymbolAndPercentage string  `json:"stockSymbolAndPercentage"`
}

type PortfolioRequest struct {
	Tradeid uint32 `json:"tradeid"`
}

type TransData struct {
	TradeId        uint32  `json:"tradeid"`
	Stocks         string  `json:"stocks"`
	UnvestedAmount float64 `json:"unvestedAmount"`
}

type PortfolioResponse struct {
	Stocks             string  `json:"stocks"`
	CurrentMarketValue float64 `json:"currentMarketValue"`
	UnvestedAmount     float64 `json:"unvestedAmount"`
}

var stockReq StockRequest
var preq PortfolioRequest

func main() {
	if len(os.Args) != 3 {

		log.Fatal(0)
	}
	//service := os.Args[1]

	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	if os.Args[1] == "buy" {
		fmt.Printf("Buy Stocks: ")
		content := []byte(os.Args[2])
		err = json.Unmarshal(content, &stockReq)
		var reply TransData
		err = client.Call("ShareMarket.BuyStock", stockReq, &reply)
		if err != nil {
			log.Fatal("error:", err)
		}
		fmt.Printf("%+v\n", reply)

	} else if os.Args[1] == "checkPortfolio" {
		fmt.Printf("Check Portfolio: ")
		content := []byte(os.Args[2])
		err = json.Unmarshal(content, &preq)
		var reply PortfolioResponse

		err = client.Call("ShareMarket.CheckPortfolio", preq, &reply)
		if err != nil {
			log.Fatal("error:", err)
		}
		fmt.Printf("%+v\n", reply)

	} else {

		fmt.Printf("Invalid Input")
	}
}
