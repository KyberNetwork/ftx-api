package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type GetDepositHistoryService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

func (s *GetDepositHistoryService) StartTime(startTime int64) *GetDepositHistoryService {
	s.startTime = &startTime
	return s
}

func (s *GetDepositHistoryService) EndTime(endTime int64) *GetDepositHistoryService {
	s.endTime = &endTime
	return s
}

type DepositHistory struct {
	Coin          string    `json:"coin"`
	Confirmations int       `json:"confirmations"`
	ConfirmedTime time.Time `json:"confirmedTime"`
	Fee           float64   `json:"fee"`
	ID            int64     `json:"id"`
	SentTime      time.Time `json:"sentTime"`
	Size          float64   `json:"size"`
	Status        string    `json:"status"`
	Time          time.Time `json:"time"`
	Txid          string    `json:"txid"`
}

type DepositHistoryResponse struct {
	basicReponse
	Result []DepositHistory `json:"result"`
}

func (s *GetDepositHistoryService) Do(ctx context.Context) ([]DepositHistory, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/deposits"), true)
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
	var result DepositHistoryResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
