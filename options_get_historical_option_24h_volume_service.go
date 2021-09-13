package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetHistorical24HOptionVolumeService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetHistorical24HOptionVolumeService) StartTime(startTime int64) *GetHistorical24HOptionVolumeService {
	s.startTime = &startTime
	return s
}

func (s *GetHistorical24HOptionVolumeService) EndTime(endTime int64) *GetHistorical24HOptionVolumeService {
	s.endTime = &endTime
	return s
}

type Historical24HVolume struct {
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	NumContracts float64   `json:"numContracts"`
}

type GetHistoricalOptionVolumeResponse struct {
	basicReponse
	Result []Historical24HVolume `json:"result"`
}

func (s *GetHistorical24HOptionVolumeService) Do(ctx context.Context) ([]Historical24HVolume, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/historical_volumes/BTC"), false)
	if s.startTime != nil {
		r.setParam("start_time", Int64ToString(*s.startTime))
	}
	if s.endTime != nil {
		r.setParam("end_time", Int64ToString(*s.endTime))
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetHistoricalOptionVolumeResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
