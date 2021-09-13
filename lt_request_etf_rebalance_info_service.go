package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type RequestETFRebalanceInfoService struct {
	c *Client
}

type ETFRebalanceInfo struct {
	OrderSizeList []string `json:"orderSizeList"`
	Side          string   `json:"side"`
	Time          string   `json:"time"`
}

type RequestETFRebalanceInfoResponse struct {
	basicReponse
	Result map[string]ETFRebalanceInfo `json:"result"`
}

func (s *RequestETFRebalanceInfoService) Do(ctx context.Context) (map[string]ETFRebalanceInfo, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/etfs/rebalance_info"), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result RequestETFRebalanceInfoResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
