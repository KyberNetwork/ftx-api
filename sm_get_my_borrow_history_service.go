package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetMyBorrowHistoryService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetMyBorrowHistoryService) StartTime(startTime int64) *GetMyBorrowHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetMyBorrowHistoryService) EndTime(endTime int64) *GetMyBorrowHistoryService {
	s.endTime = &endTime
	return s
}

type BorrowHistory struct {
	Coin string    `json:"coin"`
	Cost float64   `json:"cost"`
	Rate float64   `json:"rate"`
	Size float64   `json:"size"`
	Time time.Time `json:"time"`
}

type GetMyBorrowHistoryResponse struct {
	basicReponse
	Result []BorrowHistory `json:"result"`
}

func (s *GetMyBorrowHistoryService) Do(ctx context.Context) ([]BorrowHistory, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/spot_margin/borrow_history"), true)
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
	var result GetMyBorrowHistoryResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
