package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetTriggerOrderHistoryService struct {
	c           *Client
	market      *string
	triggerType *TriggerType
	orderSide   *OrderSide
	orderType   *OrderType
	startTime   *int64
	endTime     *int64
}

func (s *GetTriggerOrderHistoryService) Market(market string) *GetTriggerOrderHistoryService {
	s.market = &market
	return s
}

func (s *GetTriggerOrderHistoryService) TriggerType(triggerType TriggerType) *GetTriggerOrderHistoryService {
	s.triggerType = &triggerType
	return s
}

func (s *GetTriggerOrderHistoryService) Side(orderSide OrderSide) *GetTriggerOrderHistoryService {
	s.orderSide = &orderSide
	return s
}

func (s *GetTriggerOrderHistoryService) OrderType(orderType OrderType) *GetTriggerOrderHistoryService {
	s.orderType = &orderType
	return s
}

func (s *GetTriggerOrderHistoryService) StartTime(startTime int64) *GetTriggerOrderHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetTriggerOrderHistoryService) EndTime(endTime int64) *GetTriggerOrderHistoryService {
	s.endTime = &endTime
	return s
}

type TriggerOrderHistoryResponse struct {
	basicReponse
	Result      []TriggerOrder `json:"result"`
	HasMoreData bool           `json:"hasMoreData"`
}

func (s *GetTriggerOrderHistoryService) Do(ctx context.Context) ([]TriggerOrder, bool, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/conditional_orders/history"), true)
	if s.market != nil {
		r.setParam("market", *s.market)
	}
	if s.triggerType != nil {
		r.setParam("type", string(*s.triggerType))
	}
	if s.orderSide != nil {
		r.setParam("side", string(*s.orderSide))
	}
	if s.orderType != nil {
		r.setParam("orderType", string(*s.orderType))
	}
	if s.startTime != nil {
		r.setParam("start_time", Int64ToString(*s.startTime))
	}
	if s.endTime != nil {
		r.setParam("end_time", Int64ToString(*s.endTime))
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, false, err
	}
	var result TriggerOrderHistoryResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, false, err
	}
	if !result.Success {
		return nil, false, errors.New(result.Error)
	}
	return result.Result, result.HasMoreData, nil
}
