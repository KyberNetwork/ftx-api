package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type ChangeAccountLeverageService struct {
	c      *Client
	params LeverageParams
}

type LeverageParams struct {
	Leverage float64 `json:"leverage"`
}

func (s *ChangeAccountLeverageService) Params(params LeverageParams) *ChangeAccountLeverageService {
	s.params = params
	return s
}

type ChangeAccountLeverageResponse struct {
	basicReponse
}

func (s *ChangeAccountLeverageService) Do(ctx context.Context) error {
	r := newRequest(http.MethodPost, endPointWithFormat("/account/leverage"), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return err
	}
	var result ChangeAccountLeverageResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return err
	}
	if !result.Success {
		return errors.New(result.Error)
	}
	return nil
}
