package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetPublicOptionsTradesService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetPublicOptionsTradesService) StartTime(startTime int64) *GetPublicOptionsTradesService {
	s.startTime = &startTime
	return s
}

func (s *GetPublicOptionsTradesService) EndTime(endTime int64) *GetPublicOptionsTradesService {
	s.endTime = &endTime
	return s
}

type PublicOptionTrade struct {
	ID     float64   `json:"id"`
	Option Option    `json:"option"`
	Price  float64   `json:"price"`
	Size   float64   `json:"size"`
	Time   time.Time `json:"time"`
}

type GetPublicOptionsTradesResponse struct {
	basicReponse
	Result []PublicOptionTrade `json:"result"`
}

func (s *GetPublicOptionsTradesService) Do(ctx context.Context) ([]PublicOptionTrade, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/trades"), false)
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
	var result GetPublicOptionsTradesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
