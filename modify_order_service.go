package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ModifyOrderService struct {
	c       *Client
	orderID int64
	params  ModifyOrderParams
}

type ModifyOrderParams struct {
	Price    *float64 `json:"price,omitempty"`
	Size     *float64 `json:"size,omitempty"`
	ClientID *string  `json:"clientID,omitempty"`
}

func (s *ModifyOrderService) OrderID(orderID int64) *ModifyOrderService {
	s.orderID = orderID
	return s
}

func (s *ModifyOrderService) Params(params ModifyOrderParams) *ModifyOrderService {
	s.params = params
	return s
}

type ModifyOrderResponse struct {
	basicReponse
	Result *Order `json:"result"`
}

func (s *ModifyOrderService) Do(ctx context.Context) (*Order, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/orders/%d/modify", s.orderID), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result ModifyOrderResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
