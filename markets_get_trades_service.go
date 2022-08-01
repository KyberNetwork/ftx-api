package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetTradesService struct {
	c          *Client
	marketName string
	startTime  *int64
	endTime    *int64
}

func (s *GetTradesService) MarketName(marketName string) *GetTradesService {
	s.marketName = marketName
	return s
}

func (s *GetTradesService) StartTime(startTime int64) *GetTradesService {
	s.startTime = &startTime
	return s
}

func (s *GetTradesService) EndTime(endTime int64) *GetTradesService {
	s.endTime = &endTime
	return s
}

type Trade struct {
	ID          int       `json:"id"`
	Liquidation bool      `json:"liquidation"`
	Price       float64   `json:"price"`
	Side        string    `json:"side"`
	Size        float64   `json:"size"`
	Time        time.Time `json:"time"`
}

type TradesResponse struct {
	basicReponse
	Result []Trade `json:"result"`
}

func (s *GetTradesService) Do(ctx context.Context) ([]Trade, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/markets/%s/trades", s.marketName), false)
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
	var result TradesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
