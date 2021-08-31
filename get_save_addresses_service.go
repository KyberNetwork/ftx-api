package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetSaveAddressesService struct {
	c *Client
}

type SaveAddressesResponse struct {
	basicReponse
	Result []SaveAddress `json:"result"`
}

func (s *GetSaveAddressesService) Do(ctx context.Context) ([]SaveAddress, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/saved_addresses"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result SaveAddressesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
