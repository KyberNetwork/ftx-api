package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type CreateQuoteService struct {
	c         *Client
	requestID int64
	params    CreateQuoteParams
}

func (s *CreateQuoteService) RequestID(requestID int64) *CreateQuoteService {
	s.requestID = requestID
	return s
}

func (s *CreateQuoteService) Params(params CreateQuoteParams) *CreateQuoteService {
	s.params = params
	return s
}

type CreateQuoteParams struct {
	Price float64 `json:"price"`
}

type CreateQuoteResponse struct {
	basicReponse
	Result *OptionQuote `json:"result"`
}

func (s *CreateQuoteService) Do(ctx context.Context) (*OptionQuote, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/options/requests/%d/quotes", s.requestID), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result CreateQuoteResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
