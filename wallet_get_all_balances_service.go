package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetAllBalancesService struct {
	c *Client
}

type AllBalance map[string][]Balance

type AllBalancesResponse struct {
	basicReponse
	Result AllBalance `json:"result"`
}

func (s *GetAllBalancesService) Do(ctx context.Context) (AllBalance, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/all_balances"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result AllBalancesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
