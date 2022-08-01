package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type GetAirdropsService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetAirdropsService) StartTime(startTime int64) *GetAirdropsService {
	s.startTime = &startTime
	return s
}

func (s *GetAirdropsService) EndTime(endTime int64) *GetAirdropsService {
	s.endTime = &endTime
	return s
}

type Airdrop struct {
	Coin   string    `json:"coin"`
	ID     int       `json:"id"`
	Size   float64   `json:"size"`
	Time   time.Time `json:"time"`
	Status string    `json:"status"`
}

type AirdropsResponse struct {
	basicReponse
	Result []Airdrop `json:"result"`
}

func (s *GetAirdropsService) Do(ctx context.Context) ([]Airdrop, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/airdrops"), true)
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
	var result AirdropsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
