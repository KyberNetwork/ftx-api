package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Get24HOptionVolumeService struct {
	c *Client
}

type OptionVolume24H struct {
	Contracts       float64 `json:"contracts"`
	UnderlyingTotal float64 `json:"underlying_total"`
}

type Get24HOptionVolumeResponse struct {
	basicReponse
	Result *OptionVolume24H `json:"result"`
}

func (s *Get24HOptionVolumeService) Do(ctx context.Context) (*OptionVolume24H, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/stats/24h_options_volume"), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result Get24HOptionVolumeResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
