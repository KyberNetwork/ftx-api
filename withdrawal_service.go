package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type WithdrawService struct {
	c      *Client
	params WithdrawParams
}

type WithdrawParams struct {
	Coin     string  `json:"coin"`
	Size     float64 `json:"size"`
	Address  string  `json:"address"`
	Tag      *string `json:"tag,omitempty"`
	Method   *string `json:"method,omitempty"`
	Password *string `json:"password,omitempty"`
	Code     *string `json:"code,omitempty"`
}

func (s *WithdrawService) Params(params WithdrawParams) *WithdrawService {
	s.params = params
	return s
}

type WithdrawResponse struct {
	basicReponse
	Result *Withdraw `json:"result"`
}

func (s *WithdrawService) Do(ctx context.Context) (*Withdraw, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/wallet/withdrawals"), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result WithdrawResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
