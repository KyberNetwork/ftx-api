package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetSpotMarginMarketInfoService struct {
	c      *Client
	market string
}

func (s *GetSpotMarginMarketInfoService) Market(market string) *GetSpotMarginMarketInfoService {
	s.market = market
	return s
}

type CoinSpotMarginInfo struct {
	Coin          string  `json:"coin"`
	Borrowed      float64 `json:"borrowed"`
	Free          float64 `json:"free"`
	EstimatedRate float64 `json:"estimatedRate"`
	PreviousRate  float64 `json:"previousRate"`
}

type GetMarketInfoResponse struct {
	basicReponse
	Result []CoinSpotMarginInfo `json:"result"`
}

func (s *GetSpotMarginMarketInfoService) Do(ctx context.Context) ([]CoinSpotMarginInfo, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/spot_margin/market_info?market=%s", s.market), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetMarketInfoResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
