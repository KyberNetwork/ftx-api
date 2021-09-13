package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetOptionsPositionsService struct {
	c *Client
}

type OptionPosition struct {
	EntryPrice            float64  `json:"entryPrice"`
	NetSize               float64  `json:"netSize"`
	Option                Option   `json:"option"`
	Side                  Side     `json:"side"`
	Size                  float64  `json:"size"`
	PessimisticValuation  *float64 `json:"pessimisticValuation"`
	PessimisticIndexPrice *float64 `json:"pessimisticIndexPrice"`
	PessimisticVol        *float64 `json:"pessimisticVol"`
}

type GetOptionsPositionsResponse struct {
	basicReponse
	Result []OptionPosition `json:"result"`
}

func (s *GetOptionsPositionsService) Do(ctx context.Context) ([]OptionPosition, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/positions"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetOptionsPositionsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
