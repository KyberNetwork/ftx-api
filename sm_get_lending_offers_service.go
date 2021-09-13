package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetLendingOffersService struct {
	c *Client
}

type LendingOffer struct {
	Coin string  `json:"coin"`
	Rate float64 `json:"rate"`
	Size float64 `json:"size"`
}

type GetLendingOffersResponse struct {
	basicReponse
	Result []LendingOffer `json:"result"`
}

func (s *GetLendingOffersService) Do(ctx context.Context) ([]LendingOffer, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/spot_margin/offers"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetLendingOffersResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
