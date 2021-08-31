package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type CancelAllOrderService struct {
	c      *Client
	params CancelAllOrderParams
}

type CancelAllOrderParams struct {
	Market                *string    `json:"market,omitempty"`
	Side                  *OrderSide `json:"side,omitempty"`
	ConditionalOrdersOnly *bool      `json:"conditionalOrdersOnly,omitempty"`
	LimitOrdersOnly       *bool      `json:"limitOrdersOnly,omitempty"`
}

func (s *CancelAllOrderService) Params(params CancelAllOrderParams) *CancelAllOrderService {
	s.params = params
	return s
}

type CancelAllOrderResponse struct {
	basicReponse
}

func (s *CancelAllOrderService) Do(ctx context.Context) error {
	r := newRequest(http.MethodDelete, endPointWithFormat("/orders"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return err
	}
	var result CancelAllOrderResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return err
	}
	if !result.Success {
		return errors.New(result.Error)
	}
	return nil
}
