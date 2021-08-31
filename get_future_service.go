package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetFutureService struct {
	c          *Client
	futureName string
}

func (s *GetFutureService) FutureName(futureName string) *GetFutureService {
	s.futureName = futureName
	return s
}

type FutureResponse struct {
	basicReponse
	Result *Future `json:"result"`
}

func (s *GetFutureService) Do(ctx context.Context) (*Future, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/futures/%s", s.futureName), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result FutureResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
