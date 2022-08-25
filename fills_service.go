package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type FillsService struct {
	c         *Client
	market    *string
	order     *OrderBy
	orderID   *int64
	startTime *int64
	endTime   *int64
}

func (s *FillsService) Market(market string) *FillsService {
	s.market = &market
	return s
}

func (s *FillsService) Order(order OrderBy) *FillsService {
	s.order = &order
	return s
}

func (s *FillsService) OrderID(orderID int64) *FillsService {
	s.orderID = &orderID
	return s
}

func (s *FillsService) StartTime(startTime int64) *FillsService {
	s.startTime = &startTime
	return s
}

func (s *FillsService) EndTime(endTime int64) *FillsService {
	s.endTime = &endTime
	return s
}

type OrderBy string

const (
	OrderByASC  OrderBy = "asc"
	OrderByDESC OrderBy = "desc"
)

type Fill struct {
	Fee           float64     `json:"fee"`
	FeeCurrency   string      `json:"feeCurrency"`
	FeeRate       float64     `json:"feeRate"`
	Future        string      `json:"future"`
	ID            int         `json:"id"`
	Liquidity     string      `json:"liquidity"`
	Market        string      `json:"market"`
	BaseCurrency  interface{} `json:"baseCurrency"`
	QuoteCurrency interface{} `json:"quoteCurrency"`
	OrderID       int         `json:"orderId"`
	TradeID       int         `json:"tradeId"`
	Price         float64     `json:"price"`
	Side          string      `json:"side"`
	Size          float64     `json:"size"`
	Time          time.Time   `json:"time"`
	Type          string      `json:"type"`
}

type FillResponse struct {
	basicReponse
	Result []Fill `json:"result"`
}

func (s *FillsService) Do(ctx context.Context) ([]Fill, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/fills"), true)
	if s.market != nil {
		r.setParam("market", *s.market)
	}
	if s.order != nil {
		o := *s.order
		r.setParam("order", string(o))
	}
	if s.orderID != nil {
		r.setParam("orderId", int64ToString(*s.orderID))
	}
	if s.startTime != nil {
		r.setParam("start_time", int64ToString(*s.startTime))
	}
	if s.endTime != nil {
		r.setParam("end_time", int64ToString(*s.endTime))
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result FillResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
