package ftxapi

import "time"

type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
)

type Side string

const (
	SideBuy  Side = "buy"
	SideSell Side = "sell"
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

type OptionQuoteStatus string

const (
	OptionQuoteStatusOpen      OptionQuoteStatus = "open"
	OptionQuoteStatusFilled    OptionQuoteStatus = "filled"
	OptionQuoteStatusCancelled OptionQuoteStatus = "cancelled"
)

type TriggerType string

const (
	TriggerTypeStop         TriggerType = "stop"
	TriggerTypeTrailingStop TriggerType = "trailing_stop"
	TriggerTypeTakeProfit   TriggerType = "take_profit"
)

type OptionType string

const (
	OptionTypeCall OptionType = "call"
	OptionTypePut  OptionType = "put"
)

type MarketType string

const (
	SpotMarket   MarketType = "spot"
	FutureMarket MarketType = "future"
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
	Name                  string     `json:"name"`
	Enabled               bool       `json:"enabled"`
	PostOnly              bool       `json:"postOnly"`
	PriceIncrement        float64    `json:"priceIncrement"`
	SizeIncrement         float64    `json:"sizeIncrement"`
	MinProvideSize        float64    `json:"minProvideSize"`
	Last                  float64    `json:"last"`
	Bid                   float64    `json:"bid"`
	Ask                   float64    `json:"ask"`
	Price                 float64    `json:"price"`
	Type                  MarketType `json:"type"`
	FutureType            *string    `json:"futureType"`
	BaseCurrency          *string    `json:"baseCurrency"`
	IsEtfMarket           bool       `json:"isEtfMarket"`
	QuoteCurrency         *string    `json:"quoteCurrency"`
	Underlying            *string    `json:"underlying"`
	Restricted            bool       `json:"restricted"`
	HighLeverageFeeExempt bool       `json:"highLeverageFeeExempt"`
	LargeOrderThreshold   float64    `json:"largeOrderThreshold"`
	Change1H              float64    `json:"change1h"`
	Change24H             float64    `json:"change24h"`
	ChangeBod             float64    `json:"changeBod"`
	QuoteVolume24H        float64    `json:"quoteVolume24h"`
	VolumeUsd24H          float64    `json:"volumeUsd24h"`
	PriceHigh24H          float64    `json:"priceHigh24h"`
	PriceLow24H           float64    `json:"priceLow24h"`
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
	Side          Side        `json:"side"`
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
	Side             Side        `json:"side"`
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

type Basket map[string]float64

type PositionsPerShare map[string]float64

type LeveragedToken struct {
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Underlying        string            `json:"underlying"`
	Leverage          int               `json:"leverage"`
	Outstanding       float64           `json:"outstanding"`
	PricePerShare     float64           `json:"pricePerShare"`
	PositionPerShare  float64           `json:"positionPerShare"`
	PositionsPerShare PositionsPerShare `json:"positionsPerShare"`
	Basket            Basket            `json:"basket"`
	TargetComponents  []string          `json:"targetComponents"`
	UnderlyingMark    float64           `json:"underlyingMark"`
	TotalNav          float64           `json:"totalNav"`
	TotalCollateral   float64           `json:"totalCollateral"`
	ContractAddress   string            `json:"contractAddress"`
	CurrentLeverage   float64           `json:"currentLeverage"`
	Change1H          float64           `json:"change1h"`
	Change24H         float64           `json:"change24h"`
	ChangeBod         float64           `json:"changeBod"`
}

type Option struct {
	Underlying string     `json:"underlying"`
	Type       OptionType `json:"type"`
	Strike     float64    `json:"strike"`
	Expiry     time.Time  `json:"expiry"`
}

type OptionQuote struct {
	Collateral  float64           `json:"collateral"`
	ID          int64             `json:"id"`
	Option      Option            `json:"option"`
	Price       float64           `json:"price"`
	QuoteExpiry *string           `json:"quoteExpiry"`
	QuoterSide  Side              `json:"quoterSide"`
	RequestID   int64             `json:"requestId"`
	RequestSide Side              `json:"requestSide"`
	Size        float64           `json:"size"`
	Status      OptionQuoteStatus `json:"status"`
	Time        time.Time         `json:"time"`
}
