package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetBorrowRatesService struct {
	c *Client
}

type BorrowRate struct {
	Coin     string  `json:"coin"`
	Estimate float64 `json:"estimate"`
	Previous float64 `json:"previous"`
}

type GetBorrowRatesResponse struct {
	basicReponse
	Result []BorrowRate `json:"result"`
}

func (s *GetBorrowRatesService) Do(ctx context.Context) ([]BorrowRate, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/spot_margin/borrow_rates"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetBorrowRatesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
