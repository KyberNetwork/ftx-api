package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetCoinsService struct {
	c *Client
}

type Coin struct {
	Bep2Asset        *string  `json:"bep2Asset"`
	CanConvert       bool     `json:"canConvert"`
	CanDeposit       bool     `json:"canDeposit"`
	CanWithdraw      bool     `json:"canWithdraw"`
	Collateral       bool     `json:"collateral"`
	CollateralWeight float64  `json:"collateralWeight"`
	CreditTo         *string  `json:"creditTo"`
	Erc20Contract    string   `json:"erc20Contract"`
	Fiat             bool     `json:"fiat"`
	HasTag           bool     `json:"hasTag"`
	ID               string   `json:"id"`
	IsToken          bool     `json:"isToken"`
	Methods          []string `json:"methods"`
	Name             string   `json:"name"`
	SplMint          string   `json:"splMint"`
	Trc20Contract    string   `json:"trc20Contract"`
	UsdFungible      bool     `json:"usdFungible"`
}

type CoinsResponse struct {
	basicReponse
	Result []Coin `json:"result"`
}

func (s *GetCoinsService) Do(ctx context.Context) ([]Coin, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/coins"), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result CoinsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
