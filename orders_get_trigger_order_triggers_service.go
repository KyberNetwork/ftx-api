package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type GetTriggerOrderTriggersService struct {
	c                  *Client
	conditionalOrderID int64
}

type OrderTrigger struct {
	Error      error     `json:"error"`
	FilledSize *float64  `json:"filledSize"`
	OrderSize  *float64  `json:"orderSize"`
	OrderID    *int64    `json:"orderId"`
	Time       time.Time `json:"time"`
}

type TriggerOrderTriggersResponse struct {
	basicReponse
	Result []OrderTrigger `json:"result"`
}

func (s *GetTriggerOrderTriggersService) Do(ctx context.Context) ([]OrderTrigger, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/conditional_orders/%d/triggers", s.conditionalOrderID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result TriggerOrderTriggersResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
