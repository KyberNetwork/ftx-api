package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type CreateSubAccountService struct {
	c      *Client
	params CreateSubAccountParams
}

type CreateSubAccountParams struct {
	Nickname string `json:"nickname"`
}

func (s *CreateSubAccountService) Params(params CreateSubAccountParams) *CreateSubAccountService {
	s.params = params
	return s
}

type CreateSubAccountResponse struct {
	basicReponse
	Result *SubAccount `json:"result"`
}

func (s *CreateSubAccountService) Do(ctx context.Context) (*SubAccount, error) {
	r := newRequest(http.MethodPost, "/subaccounts", true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result CreateSubAccountResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
