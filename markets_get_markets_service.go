package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetMarketsService struct {
	c *Client
}

type MarketsResponse struct {
	basicReponse
	Result []Market `json:"result"`
}

func (s *GetMarketsService) Do(ctx context.Context) ([]Market, error) {
	r := newRequest(http.MethodGet, "/markets", false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result MarketsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
