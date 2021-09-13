package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetSubAccountBalanceService struct {
	c        *Client
	nickName string
}

func (s *GetSubAccountBalanceService) NickName(nickName string) *GetSubAccountBalanceService {
	s.nickName = nickName
	return s
}

type GetSubAccountBalanceResponse struct {
	basicReponse
	Result []Balance `json:"result"`
}

func (s *GetSubAccountBalanceService) Do(ctx context.Context) ([]Balance, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/subaccounts/%s/balances", s.nickName), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetSubAccountBalanceResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
