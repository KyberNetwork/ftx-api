package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetAccountOptionsInfoService struct {
	c *Client
}

type AccountOptionsInfo struct {
	UsdBalance                   float64 `json:"usdBalance"`
	LiquidationPrice             float64 `json:"liquidationPrice"`
	Liquidating                  bool    `json:"liquidating"`
	MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`
	InitialMarginRequirement     float64 `json:"initialMarginRequirement"`
}

type GetAccountOptionsInfoResponse struct {
	basicReponse
	Result *AccountOptionsInfo `json:"result"`
}

func (s *GetAccountOptionsInfoService) Do(ctx context.Context) (*AccountOptionsInfo, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/account_info"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetAccountOptionsInfoResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
