package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type ListLeveragedTokenRedemptionRequestsService struct {
	c *Client
}

type LeveragedTokenRedemptionRequest struct {
	ID          int64     `json:"id"`
	Token       string    `json:"token"`
	Size        float64   `json:"size"`
	Pending     bool      `json:"pending"`
	Price       float64   `json:"price"`
	Proceeds    float64   `json:"proceeds"`
	Fee         float64   `json:"fee"`
	RequestedAt time.Time `json:"requestedAt"`
	FulfilledAt time.Time `json:"fulfilledAt"`
}

type ListLeveragedTokenRedemptionRequestsResponse struct {
	basicReponse
	Result []LeveragedTokenCreationRequest `json:"result"`
}

func (s *ListLeveragedTokenRedemptionRequestsService) Do(ctx context.Context) ([]LeveragedTokenCreationRequest, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/lt/redemptions"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result ListLeveragedTokenRedemptionRequestsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
