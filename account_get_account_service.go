package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetAccountService struct {
	c *Client
}

type AccountResponse struct {
	basicReponse
	Result *Account `json:"result"`
}

type Account struct {
	BackstopProvider             bool       `json:"backstopProvider"`
	Collateral                   float64    `json:"collateral"`
	FreeCollateral               float64    `json:"freeCollateral"`
	InitialMarginRequirement     float64    `json:"initialMarginRequirement"`
	Leverage                     float64    `json:"leverage"`
	Liquidating                  bool       `json:"liquidating"`
	MaintenanceMarginRequirement float64    `json:"maintenanceMarginRequirement"`
	MakerFee                     float64    `json:"makerFee"`
	MarginFraction               float64    `json:"marginFraction"`
	OpenMarginFraction           float64    `json:"openMarginFraction"`
	TakerFee                     float64    `json:"takerFee"`
	TotalAccountValue            float64    `json:"totalAccountValue"`
	TotalPositionSize            float64    `json:"totalPositionSize"`
	Username                     string     `json:"username"`
	Positions                    []Position `json:"positions"`
}

func (s *GetAccountService) Do(ctx context.Context) (*Account, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/account"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result AccountResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
