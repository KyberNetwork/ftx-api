package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type GetHistoricalIndexService struct {
	c          *Client
	marketName string
	resolution int64
	startTime  *int64
	endTime    *int64
}

func (s *GetHistoricalIndexService) MarketName(marketName string) *GetHistoricalIndexService {
	s.marketName = marketName
	return s
}

func (s *GetHistoricalIndexService) Resolution(resolution int64) *GetHistoricalIndexService {
	s.resolution = resolution
	return s
}

func (s *GetHistoricalIndexService) StartTime(startTime int64) *GetHistoricalIndexService {
	s.startTime = &startTime
	return s
}

func (s *GetHistoricalIndexService) EndTime(endTime int64) *GetHistoricalIndexService {
	s.endTime = &endTime
	return s
}

type HistoricalIndex struct {
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Open      float64   `json:"open"`
	StartTime time.Time `json:"startTime"`
	Volume    *float64  `json:"volume"`
}

type HistoricalIndexResponse struct {
	basicReponse
	Result []HistoricalIndex `json:"result"`
}

func (s *GetHistoricalIndexService) Do(ctx context.Context) ([]HistoricalIndex, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/indexes/%s/candles", s.marketName), false)
	r.setParam("resolution", Int64ToString(s.resolution))
	if s.startTime != nil {
		r.setParam("start_time", Int64ToString(*s.startTime))
	}
	if s.endTime != nil {
		r.setParam("end_time", Int64ToString(*s.endTime))
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result HistoricalIndexResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
