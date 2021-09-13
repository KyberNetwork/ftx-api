package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ListLeveragedTokensService struct {
	c *Client
}

type ListLeveragedTokensResponse struct {
	basicReponse
	Result []LeveragedToken `json:"result"`
}

func (s *ListLeveragedTokensService) Do(ctx context.Context) ([]LeveragedToken, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/lt/tokens"), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result ListLeveragedTokensResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
