package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetOptionsFillsService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetOptionsFillsService) StartTime(startTime int64) *GetOptionsFillsService {
	s.startTime = &startTime
	return s
}

func (s *GetOptionsFillsService) EndTime(endTime int64) *GetOptionsFillsService {
	s.endTime = &endTime
	return s
}

type OptionFill struct {
	Fee       float64   `json:"fee"`
	FeeRate   float64   `json:"feeRate"`
	ID        int64     `json:"id"`
	Liquidity string    `json:"liquidity"`
	Option    Option    `json:"option"`
	Price     float64   `json:"price"`
	QuoteID   float64   `json:"quoteId"`
	Side      string    `json:"side"`
	Size      float64   `json:"size"`
	Time      time.Time `json:"time"`
}

type GetOptionsFillsResponse struct {
	basicReponse
	Result []OptionFill `json:"result"`
}

func (s *GetOptionsFillsService) Do(ctx context.Context) ([]OptionFill, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/fills"), true)
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
	var result GetOptionsFillsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
