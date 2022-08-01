package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type GetHistoricalPricesService struct {
	c          *Client
	marketName string
	resolution int64
	startTime  *int64
	endTime    *int64
}

func (s *GetHistoricalPricesService) MarketName(marketName string) *GetHistoricalPricesService {
	s.marketName = marketName
	return s
}

func (s *GetHistoricalPricesService) Resolution(resolution int64) *GetHistoricalPricesService {
	s.resolution = resolution
	return s
}

func (s *GetHistoricalPricesService) StartTime(startTime int64) *GetHistoricalPricesService {
	s.startTime = &startTime
	return s
}

func (s *GetHistoricalPricesService) EndTime(endTime int64) *GetHistoricalPricesService {
	s.endTime = &endTime
	return s
}

type HistoricalPrice struct {
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Open      float64   `json:"open"`
	StartTime time.Time `json:"startTime"`
	Volume    float64   `json:"volume"`
}

type HistoricalPricesResponse struct {
	basicReponse
	Result []HistoricalPrice `json:"result"`
}

func (s *GetHistoricalPricesService) Do(ctx context.Context) ([]HistoricalPrice, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/markets/%s/candles", s.marketName), false)
	r.setParam("resolution", int64ToString(s.resolution))
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
	var result HistoricalPricesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
