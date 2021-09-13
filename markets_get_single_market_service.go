package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetSingleMarketService struct {
	c          *Client
	marketName string
}

func (s *GetSingleMarketService) MarketName(marketName string) *GetSingleMarketService {
	s.marketName = marketName
	return s
}

type SingleMarketResponse struct {
	basicReponse
	Result *Market `json:"result"`
}

func (s *GetSingleMarketService) Do(ctx context.Context) (*Market, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/markets/%s", s.marketName), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result SingleMarketResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
