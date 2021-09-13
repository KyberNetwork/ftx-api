package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetOpenTriggerOrdersService struct {
	c           *Client
	market      *string
	triggerType *TriggerType
}

func (s *GetOpenTriggerOrdersService) Market(market string) *GetOpenTriggerOrdersService {
	s.market = &market
	return s
}

func (s *GetOpenTriggerOrdersService) TriggerType(triggerType TriggerType) *GetOpenTriggerOrdersService {
	s.triggerType = &triggerType
	return s
}

type OpenTriggerOrdersResponse struct {
	basicReponse
	Result []Order `json:"result"`
}

func (s *GetOpenTriggerOrdersService) Do(ctx context.Context) ([]Order, error) {
	r := newRequest(http.MethodGet, "/conditional_orders", true)
	if s.market != nil {
		r.setParam("market", *s.market)
	}
	if s.triggerType != nil {
		r.setParam("type", string(*s.triggerType))
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result OpenTriggerOrdersResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
