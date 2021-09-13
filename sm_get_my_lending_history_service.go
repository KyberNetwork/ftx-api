package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetMyLendingHistoryService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetMyLendingHistoryService) StartTime(startTime int64) *GetMyLendingHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetMyLendingHistoryService) EndTime(endTime int64) *GetMyLendingHistoryService {
	s.endTime = &endTime
	return s
}

type UserLendingHistory struct {
	Coin     string    `json:"coin"`
	Proceeds float64   `json:"proceeds"`
	Rate     float64   `json:"rate"`
	Size     float64   `json:"size"`
	Time     time.Time `json:"time"`
}

type GetMyLendingHistoryResponse struct {
	basicReponse
	Result []UserLendingHistory `json:"result"`
}

func (s *GetMyLendingHistoryService) Do(ctx context.Context) ([]UserLendingHistory, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/spot_margin/lending_history"), true)
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
	var result GetMyLendingHistoryResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
