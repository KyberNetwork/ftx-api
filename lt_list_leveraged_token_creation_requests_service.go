package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type ListLeveragedTokenCreationRequestsService struct {
	c *Client
}

type LeveragedTokenCreationRequest struct {
	ID            int64     `json:"id"`
	Token         string    `json:"token"`
	RequestedSize float64   `json:"requestedSize"`
	Pending       bool      `json:"pending"`
	CreatedSize   float64   `json:"createdSize"`
	Price         float64   `json:"price"`
	Cost          float64   `json:"cost"`
	Fee           float64   `json:"fee"`
	RequestedAt   time.Time `json:"requestedAt"`
	FulfilledAt   time.Time `json:"fulfilledAt"`
}

type ListLeveragedTokenCreationRequestsResponse struct {
	basicReponse
	Result []LeveragedTokenCreationRequest `json:"result"`
}

func (s *ListLeveragedTokenCreationRequestsService) Do(ctx context.Context) ([]LeveragedTokenCreationRequest, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/lt/creations"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result ListLeveragedTokenCreationRequestsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
