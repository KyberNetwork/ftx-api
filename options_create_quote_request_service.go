package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type CreateQuoteRequestService struct {
	c      *Client
	params CreateQuoteRequestParams
}

func (s *CreateQuoteRequestService) Params(params CreateQuoteRequestParams) *CreateQuoteRequestService {
	s.params = params
	return s
}

type CreateQuoteRequestParams struct {
	Underlying     string   `json:"underlying"`
	Type           string   `json:"type"`
	Strike         float64  `json:"strike"`
	Expiry         int64    `json:"expiry"`
	Side           Side     `json:"side"`
	Size           float64  `json:"size"`
	LimitPrice     *float64 `json:"limitPrice,omitempty"`
	HideLimitPrice bool     `json:"hideLimitPrice"`
	RequestExpiry  *float64 `json:"requestExpiry,omitempty"`
	CounterpartyId *int64   `json:"counterpartyId,omitempty"`
}

type CreateQuoteRequest struct {
	ID            int64       `json:"id"`
	Option        Option      `json:"option"`
	Expiry        time.Time   `json:"expiry"`
	Strike        float64     `json:"strike"`
	Type          OptionType  `json:"type"`
	Underlying    string      `json:"underlying"`
	RequestExpiry time.Time   `json:"requestExpiry"`
	Side          Side        `json:"side"`
	Size          float64     `json:"size"`
	Status        OrderStatus `json:"status"`
	Time          time.Time   `json:"time"`
}

type CreateQuoteRequestResponse struct {
	basicReponse
	Result *CreateQuoteRequest `json:"result"`
}

func (s *CreateQuoteRequestService) Do(ctx context.Context) (*CreateQuoteRequest, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/options/requests"), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result CreateQuoteRequestResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
