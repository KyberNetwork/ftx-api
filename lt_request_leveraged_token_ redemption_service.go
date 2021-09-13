package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type RequestLeveragedTokenRedemptionService struct {
	c         *Client
	tokenName string
	params    RequestLeveragedTokenRedemptionParams
}

type RequestLeveragedTokenRedemptionParams struct {
	Size string `json:"size"`
}

func (s *RequestLeveragedTokenRedemptionService) Params(params RequestLeveragedTokenRedemptionParams) *RequestLeveragedTokenRedemptionService {
	s.params = params
	return s
}

type RequestLeveragedTokenRedemption struct {
	ID            int64     `json:"id"`
	Token         string    `json:"token"`
	RequestedSize float64   `json:"requestedSize"`
	Cost          float64   `json:"cost"`
	Pending       bool      `json:"pending"`
	RequestedAt   time.Time `json:"requestedAt"`
}

type RequestLeveragedTokenRedemptionResponse struct {
	basicReponse
	Result *RequestLeveragedTokenRedemption `json:"result"`
}

func (s *RequestLeveragedTokenRedemptionService) Do(ctx context.Context) (*RequestLeveragedTokenRedemption, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/lt/%s/redeem", s.tokenName), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result RequestLeveragedTokenRedemptionResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
