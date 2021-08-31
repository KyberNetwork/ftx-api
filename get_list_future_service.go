package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type GetListFutureService struct {
	c *Client
}

type Future struct {
	Ask                 float64   `json:"ask"`
	Bid                 float64   `json:"bid"`
	Change1H            float64   `json:"change1h"`
	Change24H           float64   `json:"change24h"`
	ChangeBod           float64   `json:"changeBod"`
	VolumeUsd24H        float64   `json:"volumeUsd24h"`
	Volume              float64   `json:"volume"`
	Description         string    `json:"description"`
	Enabled             bool      `json:"enabled"`
	Expired             bool      `json:"expired"`
	Expiry              time.Time `json:"expiry"`
	Index               float64   `json:"index"`
	ImfFactor           float64   `json:"imfFactor"`
	Last                float64   `json:"last"`
	LowerBound          float64   `json:"lowerBound"`
	Mark                float64   `json:"mark"`
	Name                string    `json:"name"`
	OpenInterest        float64   `json:"openInterest"`
	OpenInterestUsd     float64   `json:"openInterestUsd"`
	Perpetual           bool      `json:"perpetual"`
	PositionLimitWeight float64   `json:"positionLimitWeight"`
	PostOnly            bool      `json:"postOnly"`
	PriceIncrement      float64   `json:"priceIncrement"`
	SizeIncrement       float64   `json:"sizeIncrement"`
	Underlying          string    `json:"underlying"`
	UpperBound          float64   `json:"upperBound"`
	Type                string    `json:"type"`
}

type ListFutureResponse struct {
	basicReponse
	Result []Future `json:"result"`
}

func (s *GetListFutureService) Do(ctx context.Context) ([]Future, error) {
	r := newRequest(http.MethodGet, "/futures", false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result ListFutureResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
