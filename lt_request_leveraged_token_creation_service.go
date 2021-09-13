package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type RequestLeveragedTokenCreationService struct {
	c         *Client
	tokenName string
	params    RequestLeveragedTokenCreationParams
}

type RequestLeveragedTokenCreationParams struct {
	Size string `json:"size"`
}

func (s *RequestLeveragedTokenCreationService) Params(params RequestLeveragedTokenCreationParams) *RequestLeveragedTokenCreationService {
	s.params = params
	return s
}

type RequestLeveragedTokenCreation struct {
	ID            int64     `json:"id"`
	Token         string    `json:"token"`
	RequestedSize float64   `json:"requestedSize"`
	Cost          float64   `json:"cost"`
	Pending       bool      `json:"pending"`
	RequestedAt   time.Time `json:"requestedAt"`
}

type RequestLeveragedTokenCreationResponse struct {
	basicReponse
	Result *RequestLeveragedTokenCreation `json:"result"`
}

func (s *RequestLeveragedTokenCreationService) Do(ctx context.Context) (*RequestLeveragedTokenCreation, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/lt/%s/create", s.tokenName), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result RequestLeveragedTokenCreationResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
