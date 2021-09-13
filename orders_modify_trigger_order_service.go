package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ModifyTriggerOrderService struct {
	c       *Client
	orderID int64
	params  ModifyTriggerOrderParams
}

type ModifyTriggerOrderParams struct {
	Size         float64  `json:"size"`
	TrailValue   *float64 `json:"trailValue,omitempty"`
	TriggerPrice *float64 `json:"triggerPrice,omitempty"`
	OrderPrice   *float64 `json:"orderPrice,omitempty"`
}

func (s *ModifyTriggerOrderService) OrderID(orderID int64) *ModifyTriggerOrderService {
	s.orderID = orderID
	return s
}

func (s *ModifyTriggerOrderService) Params(params ModifyTriggerOrderParams) *ModifyTriggerOrderService {
	s.params = params
	return s
}

type ModifyTriggerOrderResponse struct {
	basicReponse
	Result *TriggerOrder `json:"result"`
}

func (s *ModifyTriggerOrderService) Do(ctx context.Context) (*TriggerOrder, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/conditional_orders/%d/modify", s.orderID), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result ModifyTriggerOrderResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
