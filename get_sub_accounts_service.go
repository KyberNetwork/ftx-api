package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetSubAccountsService struct {
	c *Client
}

type SubAccount struct {
	Nickname    string `json:"nickname"`
	Deletable   bool   `json:"deletable"`
	Editable    bool   `json:"editable"`
	Competition bool   `json:"competition"`
}

type SubAccountsResponse struct {
	basicReponse
	Result []SubAccount `json:"result"`
}

func (s *GetSubAccountsService) Do(ctx context.Context) ([]SubAccount, error) {
	r := newRequest(http.MethodGet, "/subaccounts", true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result SubAccountsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
