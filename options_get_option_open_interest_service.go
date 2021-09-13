package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetOptionOpenInterestService struct {
	c *Client
}

type OptionOpenInterest struct {
	OpenInterest float64 `json:"openInterest"`
}

type GetOptionOpenInterestResponse struct {
	basicReponse
	Result *OptionOpenInterest `json:"result"`
}

func (s *GetOptionOpenInterestService) Do(ctx context.Context) (*OptionOpenInterest, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/open_interest/BTC"), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetOptionOpenInterestResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
