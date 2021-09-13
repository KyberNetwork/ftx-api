package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type CancelOrderByClientIDService struct {
	c        *Client
	clientID string
}

func (s *CancelOrderByClientIDService) ClientID(clientID string) *CancelOrderByClientIDService {
	s.clientID = clientID
	return s
}

type CancelOrderByClientIDResponse struct {
	basicReponse
}

func (s *CancelOrderByClientIDService) Do(ctx context.Context) error {
	r := newRequest(http.MethodDelete, endPointWithFormat("/orders/by_client_id/%s", s.clientID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return err
	}
	var result CancelOrderByClientIDResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return err
	}
	if !result.Success {
		return errors.New(result.Error)
	}
	return nil
}
