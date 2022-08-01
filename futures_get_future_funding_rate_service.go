package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type GetFutureFundingRateService struct {
	c         *Client
	future    *string
	startTime *int64
	endTime   *int64
}

func (s *GetFutureFundingRateService) Future(future string) *GetFutureFundingRateService {
	s.future = &future
	return s
}

func (s *GetFutureFundingRateService) StartTime(startTime int64) *GetFutureFundingRateService {
	s.startTime = &startTime
	return s
}

func (s *GetFutureFundingRateService) EndTime(endTime int64) *GetFutureFundingRateService {
	s.endTime = &endTime
	return s
}

type FutureFundingRate struct {
	Future string    `json:"future"`
	Rate   float64   `json:"rate"`
	Time   time.Time `json:"time"`
}

type FutureFundingRateResponse struct {
	basicReponse
	Result []FutureFundingRate `json:"result"`
}

func (s *GetFutureFundingRateService) Do(ctx context.Context) ([]FutureFundingRate, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/funding_rates"), false)
	if s.future != nil {
		r.setParam("future", *s.future)
	}
	if s.startTime != nil {
		r.setParam("start_time", int64ToString(*s.startTime))
	}
	if s.endTime != nil {
		r.setParam("end_time", int64ToString(*s.endTime))
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result FutureFundingRateResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
