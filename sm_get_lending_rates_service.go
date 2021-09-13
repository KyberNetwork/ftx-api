package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetLendingRatesService struct {
	c *Client
}

type LendingRate struct {
	Coin     string  `json:"coin"`
	Estimate float64 `json:"estimate"`
	Previous float64 `json:"previous"`
}

type GetLendingRatesResponse struct {
	basicReponse
	Result []LendingRate `json:"result"`
}

func (s *GetLendingRatesService) Do(ctx context.Context) ([]LendingRate, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/spot_margin/lending_rates"), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetLendingRatesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
