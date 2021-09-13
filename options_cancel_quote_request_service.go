package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type CancelQuoteRequestService struct {
	c         *Client
	requestID int64
}

func (s *CancelQuoteRequestService) RequestID(requestID int64) *CancelQuoteRequestService {
	s.requestID = requestID
	return s
}

type CancelQuoteRequest struct {
	ID            int64             `json:"id"`
	Option        Option            `json:"option"`
	RequestExpiry time.Time         `json:"requestExpiry"`
	Side          Side              `json:"side"`
	Size          float64           `json:"size"`
	Status        OptionQuoteStatus `json:"status"`
	Time          time.Time         `json:"time"`
}

type CancelQuoteRequestResponse struct {
	basicReponse
	Result *CancelQuoteRequest `json:"result"`
}

func (s *CancelQuoteRequestService) Do(ctx context.Context) (*CancelQuoteRequest, error) {
	r := newRequest(http.MethodDelete, endPointWithFormat("/options/requests/%d", s.requestID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result CancelQuoteRequestResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
