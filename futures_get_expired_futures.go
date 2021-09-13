package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type GetExpiredFuturesService struct {
	c *Client
}

type ExpiredFuture struct {
	Ask                   *float64    `json:"ask"`
	Bid                   *float64    `json:"bid"`
	Description           string      `json:"description"`
	Enabled               bool        `json:"enabled"`
	Expired               bool        `json:"expired"`
	Expiry                time.Time   `json:"expiry"`
	ExpiryDescription     string      `json:"expiryDescription"`
	Group                 string      `json:"group"`
	ImfFactor             float64     `json:"imfFactor"`
	Index                 float64     `json:"index"`
	Last                  float64     `json:"last"`
	LowerBound            float64     `json:"lowerBound"`
	MarginPrice           float64     `json:"marginPrice"`
	Mark                  float64     `json:"mark"`
	MoveStart             interface{} `json:"moveStart"`
	Name                  string      `json:"name"`
	Perpetual             bool        `json:"perpetual"`
	PositionLimitWeight   float64     `json:"positionLimitWeight"`
	PostOnly              bool        `json:"postOnly"`
	PriceIncrement        float64     `json:"priceIncrement"`
	SizeIncrement         float64     `json:"sizeIncrement"`
	Type                  string      `json:"type"`
	Underlying            string      `json:"underlying"`
	UnderlyingDescription string      `json:"underlyingDescription"`
	UpperBound            float64     `json:"upperBound"`
}

type ExpiredFuturesResponse struct {
	basicReponse
	Result []ExpiredFuture `json:"result"`
}

func (s *GetExpiredFuturesService) Do(ctx context.Context) ([]ExpiredFuture, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/expired_futures"), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result ExpiredFuturesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
