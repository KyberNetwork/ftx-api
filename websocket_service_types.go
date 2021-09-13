package ftxapi

import (
	"time"
)

type WsChannel string

const (
	// public channels
	WsChannelTicker           WsChannel = "ticker"
	WsChannelMarkets          WsChannel = "markets"
	WsChannelTrades           WsChannel = "trades"
	WsChannelOrderBook        WsChannel = "orderbook"
	WsChannelOrderbookGrouped WsChannel = "orderbookGrouped"

	// private channels
	WsChannelFills  WsChannel = "fills"
	WsChannelOrders WsChannel = "orders"
	WsChannelFTXPay WsChannel = "ftxpay"
)

type WsDataAction string

const (
	PartialWsDataAction WsDataAction = "partial"
	UpdateWsDataAction  WsDataAction = "update"
)

type Subscription struct {
	Channel  WsChannel
	Market   *string
	Grouping *int64
}

type RequestMsg struct {
	OP       string                 `json:"op"`
	Channel  *string                `json:"channel,omitempty"`
	Market   *string                `json:"market,omitempty"`
	Grouping *int64                 `json:"grouping,omitempty"`
	Args     map[string]interface{} `json:"args,omitempty"`
}

type WsReponse struct {
	OrderBookEvent        *WsOrderBookEvent
	GroupedOrderBookEvent *WsGroupedOrderBookEvent
	Ticker                *WsTickerEvent
	Markets               *WsMarketsEvent
	Trades                *WsTradesEvent
	Fills                 *WsFillsEvent
	Orders                *WsOrdersEvent
	FTXPay                *WsFTXPayEvent
}

type errorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type WsDataHandler func(res WsReponse)

type WsErrorHandler func(err error)

type baseWsEvent struct {
	Type    WsDataAction `json:"type"`
	Channel WsChannel    `json:"channel"`
	Market  string       `json:"market"`
}

type WsTickerEvent struct {
	baseWsEvent
	Data WsTicker `json:"data"`
}

type WsTicker struct {
	Ask     *float64 `json:"ask"`
	AskSize *float64 `json:"askSize"`
	Bid     *float64 `json:"bid"`
	BidSize *float64 `json:"bidSize"`
	Last    *float64 `json:"last"`
	Time    float64  `json:"time"`
}

type WsMarketsData struct {
	Action WsDataAction `json:"action"`
	Data   map[string]WsMarket
}

type WsMarketsEvent struct {
	baseWsEvent
	Data WsMarketsData `json:"data"`
}

type WsFuture struct {
	Description           string      `json:"description"`
	Enabled               bool        `json:"enabled"`
	Expired               bool        `json:"expired"`
	Expiry                *time.Time  `json:"expiry"`
	ExpiryDescription     string      `json:"expiryDescription"`
	Group                 string      `json:"group"`
	ImfFactor             float64     `json:"imfFactor"`
	MoveStart             interface{} `json:"moveStart"`
	Name                  string      `json:"name"`
	Perpetual             bool        `json:"perpetual"`
	PositionLimitWeight   int         `json:"positionLimitWeight"`
	PostOnly              bool        `json:"postOnly"`
	Type                  string      `json:"type"`
	Underlying            string      `json:"underlying"`
	UnderlyingDescription string      `json:"underlyingDescription"`
}

type WsMarket struct {
	BaseCurrency          *string   `json:"baseCurrency"`
	QuoteCurrency         *string   `json:"quoteCurrency"`
	Enabled               bool      `json:"enabled"`
	Future                *WsFuture `json:"future"`
	HighLeverageFeeExempt bool      `json:"highLeverageFeeExempt"`
	Name                  string    `json:"name"`
	PostOnly              bool      `json:"postOnly"`
	PriceIncrement        float64   `json:"priceIncrement"`
	Restricted            bool      `json:"restricted"`
	SizeIncrement         float64   `json:"sizeIncrement"`
	Type                  string    `json:"type"`
	Underlying            *string   `json:"underlying"`
}

type WsTradesEvent struct {
	baseWsEvent
	Data []WsTrade `json:"data"`
}

type WsTrade struct {
	Price       float64   `json:"price"`
	Size        float64   `json:"size"`
	Side        Side      `json:"side"`
	Liquidation bool      `json:"liquidation"`
	Time        time.Time `json:"time"`
}

type WsOrderBookEvent struct {
	baseWsEvent
	Data WsOrderBook `json:"data"`
}

type WsOrderBook struct {
	Action   WsDataAction `json:"action"`
	Asks     []Feed       `json:"asks"`
	Bids     []Feed       `json:"bids"`
	Checksum int64        `json:"checksum"`
	Time     float64      `json:"time"`
}

type WsGroupedOrderBookEvent struct {
	baseWsEvent
	Grouping float64            `json:"grouping"`
	Data     WsGroupedOrderBook `json:"data"`
}

type WsGroupedOrderBook struct {
	Asks []Feed `json:"asks"`
	Bids []Feed `json:"bids"`
}

type WsFillsEvent struct {
	baseWsEvent
	Data WsFills `json:"data"`
}

type WsFills struct {
	BaseCurrency  *string   `json:"baseCurrency"`
	Fee           float64   `json:"fee"`
	FeeCurrency   string    `json:"feeCurrency"`
	FeeRate       float64   `json:"feeRate"`
	Future        *string   `json:"future"`
	ID            int64     `json:"id"`
	Liquidity     string    `json:"liquidity"`
	Market        string    `json:"market"`
	OrderID       int64     `json:"orderId"`
	Price         float64   `json:"price"`
	QuoteCurrency string    `json:"quoteCurrency"`
	Side          Side      `json:"side"`
	Size          float64   `json:"size"`
	Time          time.Time `json:"time"`
	TradeID       int64     `json:"tradeId"`
	Type          string    `json:"type"`
}

type WsOrdersEvent struct {
	baseWsEvent
	Data WsOrders `json:"data"`
}

type WsOrders struct {
	AvgFillPrice  float64     `json:"avgFillPrice"`
	ClientID      *string     `json:"clientId"`
	CreatedAt     time.Time   `json:"createdAt"`
	FilledSize    float64     `json:"filledSize"`
	ID            int64       `json:"id"`
	Ioc           bool        `json:"ioc"`
	Liquidation   bool        `json:"liquidation"`
	Market        string      `json:"market"`
	PostOnly      bool        `json:"postOnly"`
	Price         float64     `json:"price"`
	ReduceOnly    bool        `json:"reduceOnly"`
	RemainingSize float64     `json:"remainingSize"`
	Side          Side        `json:"side"`
	Size          float64     `json:"size"`
	Status        OrderStatus `json:"status"`
	Type          OrderType   `json:"type"`
}

type WsFTXPayEvent struct {
	baseWsEvent
	Data WsFTXPay `json:"data"`
}

type WsFTXPay struct {
	App     string `json:"app"`
	Payment string `json:"payment"`
	Status  string `json:"status"`
}
