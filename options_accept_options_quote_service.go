package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AcceptOptionsQuoteService struct {
	c       *Client
	quoteID int64
}

func (s *AcceptOptionsQuoteService) QuoteID(quoteID int64) *AcceptOptionsQuoteService {
	s.quoteID = quoteID
	return s
}

type AcceptOptionsQuoteResponse struct {
	basicReponse
	Result *OptionQuote `json:"result"`
}

func (s *AcceptOptionsQuoteService) Do(ctx context.Context) (*OptionQuote, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/options/quotes/%d/accept", s.quoteID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result AcceptOptionsQuoteResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
