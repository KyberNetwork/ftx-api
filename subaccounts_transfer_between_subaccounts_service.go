package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type TransferBetweenSubAccountsService struct {
	c      *Client
	params TransferBetweenSubAccountsParams
}

type TransferBetweenSubAccountsParams struct {
	Coin        string  `json:"coin"`
	Size        float64 `json:"size"`
	Source      *string `json:"source"`      // nil or "main" for main account
	Destination *string `json:"destination"` // nil or "main" for main account
}

func (s *TransferBetweenSubAccountsService) Params(params TransferBetweenSubAccountsParams) *TransferBetweenSubAccountsService {
	s.params = params
	return s
}

type TransferBetweenSubAccounts struct {
	ID     int64     `json:"id"`
	Coin   string    `json:"coin"`
	Size   float64   `json:"size"`
	Time   time.Time `json:"time"`
	Notes  string    `json:"notes"`
	Status string    `json:"status"`
}

type TransferBetweenSubAccountsResponse struct {
	basicReponse
	Result *TransferBetweenSubAccounts `json:"result"`
}

func (s *TransferBetweenSubAccountsService) Do(ctx context.Context) (*TransferBetweenSubAccounts, error) {
	r := newRequest(http.MethodPost, "/subaccounts/transfer", true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result TransferBetweenSubAccountsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
