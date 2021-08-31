package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type PlaceTriggerOrderService struct {
	c      *Client
	params PlaceTriggerOrderParams
}

type PlaceTriggerOrderParams struct {
	Market           string      `json:"market"`
	Side             OrderSide   `json:"side"`
	Size             float64     `json:"size"`
	Type             TriggerType `json:"type"`
	TrailValue       *float64    `json:"trailValue,omitempty"`
	TriggerPrice     *float64    `json:"triggerPrice,omitempty"`
	OrderPrice       *float64    `json:"orderPrice,omitempty"`
	ReduceOnly       *bool       `json:"reduceOnly,omitempty"`
	RetryUntilFilled *bool       `json:"retryUntilFilled,omitempty"`
}

func (s *PlaceTriggerOrderService) Params(params PlaceTriggerOrderParams) *PlaceTriggerOrderService {
	s.params = params
	return s
}

type PlaceTriggerOrderResponse struct {
	basicReponse
	Result *TriggerOrder `json:"result"`
}

func (s *PlaceTriggerOrderService) Do(ctx context.Context) (*TriggerOrder, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/conditional_orders"), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result PlaceTriggerOrderResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
