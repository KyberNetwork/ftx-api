package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetWithdrawHistoryService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetWithdrawHistoryService) StartTime(startTime int64) *GetWithdrawHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetWithdrawHistoryService) EndTime(endTime int64) *GetWithdrawHistoryService {
	s.endTime = &endTime
	return s
}

type WithdrawHistoryResponse struct {
	basicReponse
	Result []Withdraw `json:"result"`
}

func (s *GetWithdrawHistoryService) Do(ctx context.Context) ([]Withdraw, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/withdrawals"), true)
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
	var result WithdrawHistoryResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
