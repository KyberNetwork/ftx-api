package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetDepositAddressService struct {
	c      *Client
	coin   string
	method *string
}

func (s *GetDepositAddressService) Coin(c string) *GetDepositAddressService {
	s.coin = c
	return s
}

func (s *GetDepositAddressService) Method(method string) *GetDepositAddressService {
	s.method = &method
	return s
}

type DepositAddress struct {
	Address string  `json:"address"`
	Tag     *string `json:"tag"`
}

type DepositAddressResponse struct {
	basicReponse
	Result *DepositAddress `json:"result"`
}

func (s *GetDepositAddressService) Do(ctx context.Context) (*DepositAddress, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/deposit_address/%s", s.coin), true)
	if s.method != nil {
		r.setParam("method", *s.method)
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result DepositAddressResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
