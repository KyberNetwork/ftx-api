package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetQuotesForYourQuoteRequestService struct {
	c         *Client
	requestID int64
}

func (s *GetQuotesForYourQuoteRequestService) RequestID(requestID int64) *GetQuotesForYourQuoteRequestService {
	s.requestID = requestID
	return s
}

type QuotesForYourQuoteRequest struct {
	Collateral  float64           `json:"collateral"`
	ID          int64             `json:"id"` // quote id
	Option      Option            `json:"option"`
	Price       float64           `json:"price"`
	QuoteExpiry *string           `json:"quoteExpiry"`
	QuoterSide  Side              `json:"quoterSide"`
	RequestID   int64             `json:"requestId"` // quote request id
	RequestSide Side              `json:"requestSide"`
	Size        float64           `json:"size"`
	Status      OptionQuoteStatus `json:"status"`
	Time        time.Time         `json:"time"`
}

type GetQuotesForYourQuoteRequestResponse struct {
	basicReponse
	Result []YourQuoteRequest `json:"result"`
}

func (s *GetQuotesForYourQuoteRequestService) Do(ctx context.Context) ([]YourQuoteRequest, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/requests/%d/quotes", s.requestID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetQuotesForYourQuoteRequestResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
