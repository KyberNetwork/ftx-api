package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type ListQuoteRequestsService struct {
	c *Client
}

type QuoteRequest struct {
	ID            int64             `json:"id"`
	Option        Option            `json:"option"`
	Side          Side              `json:"side"`
	Size          float64           `json:"size"`
	Time          time.Time         `json:"time"`
	RequestExpiry time.Time         `json:"requestExpiry"`
	Status        OptionQuoteStatus `json:"status"`
	LimitPrice    *float64          `json:"limitPrice"`
}

type ListQuoteRequestsResponse struct {
	basicReponse
	Result []QuoteRequest `json:"result"`
}

func (s *ListQuoteRequestsService) Do(ctx context.Context) ([]QuoteRequest, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/requests"), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result ListQuoteRequestsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
