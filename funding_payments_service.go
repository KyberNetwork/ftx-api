package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type FundingPaymentsService struct {
	c         *Client
	future    *string
	startTime *int64
	endTime   *int64
}

func (s *FundingPaymentsService) Future(future string) *FundingPaymentsService {
	s.future = &future
	return s
}

func (s *FundingPaymentsService) StartTime(startTime int64) *FundingPaymentsService {
	s.startTime = &startTime
	return s
}

func (s *FundingPaymentsService) EndTime(endTime int64) *FundingPaymentsService {
	s.endTime = &endTime
	return s
}

type FundingPayment struct {
	Future  string    `json:"future"`
	ID      int       `json:"id"`
	Payment float64   `json:"payment"`
	Time    time.Time `json:"time"`
	Rate    float64   `json:"rate"`
}

type FundingPaymentsResponse struct {
	basicReponse
	Result []FundingPayment `json:"result"`
}

func (s *FundingPaymentsService) Do(ctx context.Context) ([]FundingPayment, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/funding_payments"), true)
	if s.future != nil {
		r.setParam("future", *s.future)
	}
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
	var result FundingPaymentsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
