package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetDepositAddressListService struct {
	c *Client
}

type DepositAddressList struct {
	Coin    string  `json:"coin"`
	Address string  `json:"address"`
	Tag     *string `json:"tag"`
}

type GetDepositAddressListResponse struct {
	basicReponse
	Result *DepositAddressList `json:"result"`
}

func (s *GetDepositAddressListService) Do(ctx context.Context) (*DepositAddressList, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/deposit_address/list"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetDepositAddressListResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
