package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type ModifyOrderByClientIDService struct {
	c        *Client
	clientID string
	params   ModifyOrderByClientIDParams
}

type ModifyOrderByClientIDParams struct {
	Price *float64 `json:"price,omitempty"`
	Size  *float64 `json:"size,omitempty"`
}

func (s *ModifyOrderByClientIDService) ClientID(clientID string) *ModifyOrderByClientIDService {
	s.clientID = clientID
	return s
}

func (s *ModifyOrderByClientIDService) Params(params ModifyOrderByClientIDParams) *ModifyOrderByClientIDService {
	s.params = params
	return s
}

type ModifyOrderByClientIDResponse struct {
	basicReponse
	Result *Order `json:"result"`
}

func (s *ModifyOrderByClientIDService) Do(ctx context.Context) (*Order, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/orders/by_client_id/%s/modify", s.clientID), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result ModifyOrderByClientIDResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
