package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetHistoricalOpenInterestService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetHistoricalOpenInterestService) StartTime(startTime int64) *GetHistoricalOpenInterestService {
	s.startTime = &startTime
	return s
}

func (s *GetHistoricalOpenInterestService) EndTime(endTime int64) *GetHistoricalOpenInterestService {
	s.endTime = &endTime
	return s
}

type HistoricalOpenInterest struct {
	Time         time.Time `json:"time"`
	NumContracts float64   `json:"numContracts"`
}

type GetHistoricalOpenInterestResponse struct {
	basicReponse
	Result []HistoricalOpenInterest `json:"result"`
}

func (s *GetHistoricalOpenInterestService) Do(ctx context.Context) ([]HistoricalOpenInterest, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/historical_open_interest/BTC"), false)
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
	var result GetHistoricalOpenInterestResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
