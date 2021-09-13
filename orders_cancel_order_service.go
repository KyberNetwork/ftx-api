package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type CancelOrderService struct {
	c       *Client
	orderID int64
}

func (s *CancelOrderService) OrderID(orderID int64) *CancelOrderService {
	s.orderID = orderID
	return s
}

type CancelOrderResponse struct {
	basicReponse
}

func (s *CancelOrderService) Do(ctx context.Context) error {
	r := newRequest(http.MethodDelete, endPointWithFormat("/orders/%d", s.orderID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return err
	}
	var result CancelOrderResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return err
	}
	if !result.Success {
		return errors.New(result.Error)
	}
	return nil
}
