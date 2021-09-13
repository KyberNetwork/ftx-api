package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetOrderStatusService struct {
	c       *Client
	orderID int64
}

func (s *GetOrderStatusService) OrderID(orderID int64) *GetOrderStatusService {
	s.orderID = orderID
	return s
}

type GetOrderStatusResponse struct {
	basicReponse
	Result *Order `json:"result"`
}

func (s *GetOrderStatusService) Do(ctx context.Context) (*Order, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/orders/%d", s.orderID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetOrderStatusResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
