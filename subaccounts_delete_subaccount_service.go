package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type DeleteSubAccountService struct {
	c      *Client
	params DeleteSubAccountParams
}

type DeleteSubAccountParams struct {
	Nickname string `json:"nickname"`
}

func (s *DeleteSubAccountService) Params(params DeleteSubAccountParams) *DeleteSubAccountService {
	s.params = params
	return s
}

type DeleteSubAccountResponse struct {
	basicReponse
}

func (s *DeleteSubAccountService) Do(ctx context.Context) error {
	r := newRequest(http.MethodDelete, "/subaccounts", true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return err
	}
	var result DeleteSubAccountResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return err
	}
	if !result.Success {
		return errors.New(result.Error)
	}
	return nil
}
