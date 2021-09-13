package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetLendingInfoService struct {
	c *Client
}

type LendingInfo struct {
	Coin     string  `json:"coin"`
	Lendable float64 `json:"lendable"`
	Locked   float64 `json:"locked"`
	MinRate  float64 `json:"minRate"`
	Offered  float64 `json:"offered"`
}

type GetLendingInfoResponse struct {
	basicReponse
	Result []LendingInfo `json:"result"`
}

func (s *GetLendingInfoService) Do(ctx context.Context) ([]LendingInfo, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/spot_margin/lending_info"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetLendingInfoResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
