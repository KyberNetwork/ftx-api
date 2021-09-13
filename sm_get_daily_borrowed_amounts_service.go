package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetDailyBorrowedAmountsService struct {
	c *Client
}

type DailyBorrowedAmounts struct {
	Coin string  `json:"coin"`
	Size float64 `json:"size"`
}

type GetDailyBorrowedAmountsResponse struct {
	basicReponse
	Result []DailyBorrowedAmounts `json:"result"`
}

func (s *GetDailyBorrowedAmountsService) Do(ctx context.Context) ([]DailyBorrowedAmounts, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/spot_margin/borrow_summary"), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetDailyBorrowedAmountsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
