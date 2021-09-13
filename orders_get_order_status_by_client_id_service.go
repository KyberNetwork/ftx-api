package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetOrderStatusByClientIDService struct {
	c        *Client
	clientID string
}

func (s *GetOrderStatusByClientIDService) ClientID(clientID string) *GetOrderStatusByClientIDService {
	s.clientID = clientID
	return s
}

type GetOrderStatusByClientIDResponse struct {
	basicReponse
	Result *Order `json:"result"`
}

func (s *GetOrderStatusByClientIDService) Do(ctx context.Context) (*Order, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/orders/by_client_id/%s", s.clientID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetOrderStatusByClientIDResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
