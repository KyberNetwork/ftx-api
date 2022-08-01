package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetLendingHistoryService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetLendingHistoryService) StartTime(startTime int64) *GetLendingHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetLendingHistoryService) EndTime(endTime int64) *GetLendingHistoryService {
	s.endTime = &endTime
	return s
}

type LendingHistory struct {
	Coin string    `json:"coin"`
	Time time.Time `json:"time"`
	Rate float64   `json:"rate"`
	Size float64   `json:"size"`
}

type GetLendingHistoryResponse struct {
	basicReponse
	Result []LendingHistory `json:"result"`
}

func (s *GetLendingHistoryService) Do(ctx context.Context) ([]LendingHistory, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/spot_margin/history"), false)
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
	var result GetLendingHistoryResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
