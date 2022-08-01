package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetWithdrawalFeesService struct {
	c       *Client
	coin    string
	size    float64
	address string
	tag     *string
}

func (s *GetWithdrawalFeesService) Coin(coin string) *GetWithdrawalFeesService {
	s.coin = coin
	return s
}

func (s *GetWithdrawalFeesService) Size(size float64) *GetWithdrawalFeesService {
	s.size = size
	return s
}

func (s *GetWithdrawalFeesService) Address(address string) *GetWithdrawalFeesService {
	s.address = address
	return s
}

func (s *GetWithdrawalFeesService) Tag(tag string) *GetWithdrawalFeesService {
	s.tag = &tag
	return s
}

type WithdrawalFee struct {
	Method    string  `json:"method"`
	Address   string  `json:"address"`
	Fee       float64 `json:"fee"`
	Congested bool    `json:"congested"`
}

type WithdrawalFeesResponse struct {
	basicReponse
	Result *WithdrawalFee `json:"result"`
}

func (s *GetWithdrawalFeesService) Do(ctx context.Context) (*WithdrawalFee, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/withdrawal_fee"), true)
	r.setParam("coin", s.coin)
	r.setParam("size", float64ToString(s.size))
	r.setParam("address", s.address)
	if s.tag != nil {
		r.setParam("tag", *s.tag)
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result WithdrawalFeesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
