package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type YourQuoteRequestsService struct {
	c *Client
}

type YourQuoteRequest struct {
	ID             int64             `json:"id"`
	Option         Option            `json:"option"`
	Side           Side              `json:"side"`
	Size           float64           `json:"size"`
	Time           time.Time         `json:"time"`
	RequestExpiry  time.Time         `json:"requestExpiry"`
	Status         OptionQuoteStatus `json:"status"`
	HideLimitPrice bool              `json:"hideLimitPrice"`
	LimitPrice     float64           `json:"limitPrice"`
	Quotes         []YourQuote       `json:"quotes"`
}

type YourQuote struct {
	Collateral  float64           `json:"collateral"`
	ID          int64             `json:"id"`
	Price       float64           `json:"price"`
	QuoteExpiry *string           `json:"quoteExpiry"`
	Status      OptionQuoteStatus `json:"status"`
	Time        time.Time         `json:"time"`
}

type YourQuoteRequestsResponse struct {
	basicReponse
	Result []YourQuoteRequest `json:"result"`
}

func (s *YourQuoteRequestsService) Do(ctx context.Context) ([]YourQuoteRequest, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/my_requests"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result YourQuoteRequestsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
