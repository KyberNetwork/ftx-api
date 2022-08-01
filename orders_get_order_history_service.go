package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetOrderHistoryService struct {
	c         *Client
	market    *string
	startTime *int64
	endTime   *int64
}

func (s *GetOrderHistoryService) Market(market string) *GetOrderHistoryService {
	s.market = &market
	return s
}

func (s *GetOrderHistoryService) StartTime(startTime int64) *GetOrderHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetOrderHistoryService) EndTime(endTime int64) *GetOrderHistoryService {
	s.endTime = &endTime
	return s
}

type OrderHistoryResponse struct {
	basicReponse
	Result      []Order `json:"result"`
	HasMoreData bool    `json:"hasMoreData"`
}

func (s *GetOrderHistoryService) Do(ctx context.Context) ([]Order, bool, error) {
	r := newRequest(http.MethodGet, "/orders/history", true)
	if s.market != nil {
		r.setParam("market", *s.market)
	}
	if s.startTime != nil {
		r.setParam("start_time", int64ToString(*s.startTime))
	}
	if s.endTime != nil {
		r.setParam("end_time", int64ToString(*s.endTime))
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, false, err
	}
	var result OrderHistoryResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, false, err
	}
	if !result.Success {
		return nil, false, errors.New(result.Error)
	}
	return result.Result, result.HasMoreData, nil
}
