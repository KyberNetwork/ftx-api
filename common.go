package ftxapi

import "time"

type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
)

type OrderSide string

const (
	OrderSideBuy  OrderSide = "buy"
	OrderSideSell OrderSide = "sell"
)

type OrderStatus string

const (
	OrderStatusNew       OrderStatus = "new"
	OrderStatusOpen      OrderStatus = "open"
	OrderStatusFilled    OrderStatus = "filled"
	OrderStatusCancelled OrderStatus = "cancelled"
	OrderStatusClosed    OrderStatus = "closed"
	OrderStatusTriggered OrderStatus = "triggered"
)

type TriggerType string

const (
	TriggerTypeStop         TriggerType = "stop"
	TriggerTypeTrailingStop TriggerType = "trailing_stop"
	TriggerTypeTakeProfit   TriggerType = "take_profit"
)

type basicReponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type Position struct {
	Cost                         float64 `json:"cost"`
	CumulativeBuySize            float64 `json:"cumulativeBuySize"`
	CumulativeSellSize           float64 `json:"cumulativeSellSize"`
	EntryPrice                   float64 `json:"entryPrice"`
	EstimatedLiquidationPrice    float64 `json:"estimatedLiquidationPrice"`
	Future                       string  `json:"future"`
	InitialMarginRequirement     float64 `json:"initialMarginRequirement"`
	LongOrderSize                float64 `json:"longOrderSize"`
	MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`
	NetSize                      float64 `json:"netSize"`
	OpenSize                     float64 `json:"openSize"`
	RealizedPnl                  float64 `json:"realizedPnl"`
	RecentAverageOpenPrice       float64 `json:"recentAverageOpenPrice"`
	RecentBreakEvenPrice         float64 `json:"recentBreakEvenPrice"`
	RecentPnl                    float64 `json:"recentPnl"`
	ShortOrderSize               float64 `json:"shortOrderSize"`
	Side                         string  `json:"side"`
	Size                         float64 `json:"size"`
	UnrealizedPnl                int     `json:"unrealizedPnl"`
	CollateralUsed               float64 `json:"collateralUsed"`
}

type Market struct {
	Name                  string  `json:"name"`
	BaseCurrency          *string `json:"baseCurrency"`
	QuoteCurrency         *string `json:"quoteCurrency"`
	QuoteVolume24H        float64 `json:"quoteVolume24h"`
	Change1H              float64 `json:"change1h"`
	Change24H             float64 `json:"change24h"`
	ChangeBod             float64 `json:"changeBod"`
	HighLeverageFeeExempt bool    `json:"highLeverageFeeExempt"`
	MinProvideSize        float64 `json:"minProvideSize"`
	Type                  string  `json:"type"`
	Underlying            string  `json:"underlying"`
	Enabled               bool    `json:"enabled"`
	Ask                   float64 `json:"ask"`
	Bid                   float64 `json:"bid"`
	Last                  float64 `json:"last"`
	PostOnly              bool    `json:"postOnly"`
	Price                 float64 `json:"price"`
	PriceIncrement        float64 `json:"priceIncrement"`
	SizeIncrement         float64 `json:"sizeIncrement"`
	Restricted            bool    `json:"restricted"`
	VolumeUsd24H          float64 `json:"volumeUsd24h"`
}

type Balance struct {
	Coin                   string  `json:"coin"`
	Free                   float64 `json:"free"`
	SpotBorrow             float64 `json:"spotBorrow"`
	Total                  float64 `json:"total"`
	UsdValue               float64 `json:"usdValue"`
	AvailableWithoutBorrow float64 `json:"availableWithoutBorrow"`
}

type Withdraw struct {
	Coin    string    `json:"coin"`
	Address string    `json:"address"`
	Tag     *string   `json:"tag"`
	Fee     float64   `json:"fee"`
	ID      int64     `json:"id"`
	Size    float64   `json:"size"`
	Status  string    `json:"status"`
	Time    time.Time `json:"time"`
	Txid    string    `json:"txid"`
}

type SaveAddress struct {
	Address          string     `json:"address"`
	Coin             string     `json:"coin"`
	Fiat             bool       `json:"fiat"`
	ID               int64      `json:"id"`
	IsPrimetrust     bool       `json:"isPrimetrust"`
	LastUsedAt       time.Time  `json:"lastUsedAt"`
	Name             string     `json:"name"`
	Tag              *string    `json:"tag"`
	Whitelisted      *bool      `json:"whitelisted"`
	WhitelistedAfter *time.Time `json:"whitelistedAfter"`
}

type Order struct {
	CreatedAt     time.Time   `json:"createdAt"`
	FilledSize    float64     `json:"filledSize"`
	Future        string      `json:"future"`
	ID            int64       `json:"id"`
	Market        string      `json:"market"`
	Price         float64     `json:"price"`
	AvgFillPrice  float64     `json:"avgFillPrice"`
	RemainingSize float64     `json:"remainingSize"`
	Side          OrderSide   `json:"side"`
	Size          float64     `json:"size"`
	Status        OrderStatus `json:"status"`
	Type          OrderType   `json:"type"`
	ReduceOnly    bool        `json:"reduceOnly"`
	Ioc           bool        `json:"ioc"`
	PostOnly      bool        `json:"postOnly"`
	ClientID      *string     `json:"clientId"`
}

type TriggerOrder struct {
	CreatedAt        time.Time   `json:"createdAt"`
	Future           string      `json:"future"`
	ID               int         `json:"id"`
	Market           string      `json:"market"`
	OrderPrice       *float64    `json:"orderPrice"`
	ReduceOnly       bool        `json:"reduceOnly"`
	Side             OrderSide   `json:"side"`
	Size             float64     `json:"size"`
	Status           OrderStatus `json:"status"`
	TrailStart       *float64    `json:"trailStart"`
	TrailValue       *float64    `json:"trailValue"`
	TriggerPrice     float64     `json:"triggerPrice"`
	TriggeredAt      *string     `json:"triggeredAt"`
	Type             TriggerType `json:"type"`
	OrderType        OrderType   `json:"orderType"`
	FilledSize       float64     `json:"filledSize"`
	AvgFillPrice     *float64    `json:"avgFillPrice"`
	RetryUntilFilled bool        `json:"retryUntilFilled"`
}