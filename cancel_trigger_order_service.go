package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type CancelTriggerOrderService struct {
	c       *Client
	orderID int64
}

func (s *CancelTriggerOrderService) OrderID(orderID int64) *CancelTriggerOrderService {
	s.orderID = orderID
	return s
}

type CancelTriggerOrderResponse struct {
	basicReponse
}

func (s *CancelTriggerOrderService) Do(ctx context.Context) error {
	r := newRequest(http.MethodDelete, endPointWithFormat("/conditional_orders/%d", s.orderID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return err
	}
	var result CancelTriggerOrderResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return err
	}
	if !result.Success {
		return errors.New(result.Error)
	}
	return nil
}
