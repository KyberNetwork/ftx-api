package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type PlaceOrderService struct {
	c      *Client
	params PlaceOrderParams
}

type PlaceOrderParams struct {
	Market            string    `json:"market"`
	Side              OrderSide `json:"side"`
	Price             float64   `json:"price"`
	Type              OrderType `json:"type"`
	Size              float64   `json:"size"`
	ReduceOnly        *bool     `json:"reduceOnly,omitempty"`
	Ioc               *bool     `json:"ioc,omitempty"`
	PostOnly          *bool     `json:"postOnly,omitempty"`
	ClientID          *string   `json:"clientId,omitempty"`
	RejectOnPriceBand *bool     `json:"rejectOnPriceBand,omitempty"`
}

func (s *PlaceOrderService) Params(params PlaceOrderParams) *PlaceOrderService {
	s.params = params
	return s
}

type PlaceOrderResponse struct {
	basicReponse
	Result *Order `json:"result"`
}

func (s *PlaceOrderService) Do(ctx context.Context) (*Order, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/orders"), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result PlaceOrderResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
