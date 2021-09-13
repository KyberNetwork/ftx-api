package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetOpenOrdersService struct {
	c      *Client
	market *string
}

func (s *GetOpenOrdersService) Market(market string) *GetOpenOrdersService {
	s.market = &market
	return s
}

type OrdersResponse struct {
	basicReponse
	Result []Order `json:"result"`
}

func (s *GetOpenOrdersService) Do(ctx context.Context) ([]Order, error) {
	r := newRequest(http.MethodGet, "/orders", true)
	if s.market != nil {
		r.setParam("market", *s.market)
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result OrdersResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
