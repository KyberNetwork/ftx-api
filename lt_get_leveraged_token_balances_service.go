package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetLeveragedTokenBalancesService struct {
	c *Client
}

type LeveragedTokenBalance struct {
	Token   string  `json:"token"`
	Balance float64 `json:"balance"`
}

type GetLeveragedTokenBalancesResponse struct {
	basicReponse
	Result []LeveragedTokenBalance `json:"result"`
}

func (s *GetLeveragedTokenBalancesService) Do(ctx context.Context) ([]LeveragedTokenBalance, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/lt/balances"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetLeveragedTokenBalancesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
