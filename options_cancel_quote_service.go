package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type CancelQuoteService struct {
	c       *Client
	quoteID int64
}

func (s *CancelQuoteService) QuoteID(quoteID int64) *CancelQuoteService {
	s.quoteID = quoteID
	return s
}

type CancelQuoteResponse struct {
	basicReponse
	Result *OptionQuote `json:"result"`
}

func (s *CancelQuoteService) Do(ctx context.Context) (*OptionQuote, error) {
	r := newRequest(http.MethodDelete, endPointWithFormat("/options/quotes/%d", s.quoteID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result CancelQuoteResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
