# ftx-api

REST & WebSocket APIs for FTX exchange

[![GoDoc](https://pkg.go.dev/github.com/KyberNetwork/ftx-api?status.svg)](https://pkg.go.dev/github.com/KyberNetwork/ftx-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/KyberNetwork/ftx-api)](https://goreportcard.com/report/github.com/KyberNetwork/ftx-api)

## How to use

```golang
package main

import (
	"context"
	"fmt"
	"time"

	ftxapi "github.com/KyberNetwork/ftx-api"
	"go.uber.org/zap"
)

var sugar = zap.NewExample().Sugar()

func main() {
	l, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(l)
	sugar = l.Sugar()

	c := ftxapi.NewClient("your api key", "your api secret", ftxapi.RestAPIEndpoint, sugar)
	futureStat, err := c.NewGetFutureStatsService().FutureName("1INCH-PERP").Do(context.Background())
	if err != nil {
		sugar.Errorw("err", "err", err)
		return
	}
	sugar.Infow("data", "futureStat", futureStat)

	res, err := c.NewPlaceOrderService().Params(ftxapi.PlaceOrderParams{
		Market: "SOL/USDT",
		Side:   ftxapi.SideBuy,
		Price:  150,
		Type:   ftxapi.OrderTypeLimit,
		Size:   0.1,
	}).Do(context.Background())

	if err != nil {
		sugar.Errorw("err", "err", err)
		return
	}

	sugar.Infow("data", "res", res)

	err = c.NewCancelOrderService().OrderID(res.ID).Do(context.Background())
	sugar.Infow("err cancel order", "err", err)

	s := ftxapi.NewWebsocketService("your api key", "your api secret", ftxapi.WebsocketEndpoint, sugar).AutoReconnect()
	err = s.Connect(handler, errHandler)
	if err != nil {
		sugar.Errorw("err", "err", err)
		return
	}
	err = s.Subscribe(ftxapi.Subscription{
		Channel: ftxapi.WsChannelOrders,
	})
	if err != nil {
		sugar.Errorw("err sub", "err", err)
		return
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func handler(res ftxapi.WsReponse) {
	sugar.Infow("data", "data", res)
}

func errHandler(err error) {
	sugar.Errorw("err", "err", err)
}

```
