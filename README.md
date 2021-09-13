# ftx-api

## REST api

```golang
package main

import (
	"context"
	"fmt"

	ftxapi "github.com/KyberNetwork/ftx-api"
	"go.uber.org/zap"
)

var sugar = zap.NewExample().Sugar()

func main() {
	l, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(l)
	sugar = l.Sugar()

	c := ftxapi.NewClient("your api key", "your api secret", ftxapi.RestAPIEndpoint, l.Sugar())
	futureStat, err := c.NewGetFutureStatsService().FutureName("1INCH-PERP").Do(context.Background())
  if err != nil {
		sugar.Errorw("err", "err", err)
		return
	}
	sugar.Infow("data", "futureStat")

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

  s := ftxapi.NewWebsocketService("your api key", "your api secret", l.Sugar()).AutoReconnect()
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
  for {}
```