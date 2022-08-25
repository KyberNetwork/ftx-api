# ftx-api

REST & WebSocket APIs for FTX exchange

[![GoDoc](https://pkg.go.dev/badge/https:/github.com/KyberNetwork/ftx-api?utm_source=godoc)](https://pkg.go.dev/github.com/KyberNetwork/ftx-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/KyberNetwork/ftx-api)](https://goreportcard.com/report/github.com/KyberNetwork/ftx-api)

## How to use

```golang
package main

import (
	"context"
	"log"
	"time"

	ftxapi "github.com/KyberNetwork/ftx-api"
)

func main() {
	c := ftxapi.NewClient("your api key", "your api secret")
	futureStat, err := c.NewGetFutureStatsService().FutureName("1INCH-PERP").Do(context.Background())
	if err != nil {
		log.Println("error get future stats", err)
		return
	}
	log.Println("future stat", futureStat)

	res, err := c.NewPlaceOrderService().Params(ftxapi.PlaceOrderParams{
		Market: "SOL/USDT",
		Side:   ftxapi.SideBuy,
		Price:  150,
		Type:   ftxapi.OrderTypeLimit,
		Size:   0.1,
	}).Do(context.Background())

	if err != nil {
		log.Println("error place order", err)
		return
	}

	log.Println("response", res)

	err = c.NewCancelOrderService().OrderID(res.ID).Do(context.Background())
	log.Println("error cancel order", err)

	s := ftxapi.NewWebsocketService("your api key", "your api secret", /*autoReconnect*/ true)
	if err := s.Connect(handler, errHandler); err != nil {
		log.Println("error connect ws", err)
		return
	}
	if err = s.Subscribe(ftxapi.Subscription{
		Channel: ftxapi.WsChannelOrders,
	}); err != nil {
		log.Println("error subscribe", err)
		return
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func handler(res ftxapi.WsReponse) {
	log.Printf("data %+v\n", res)
}

func errHandler(err error) {
	log.Printf("err = %+s\n", err)
}

```
