package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type DeleteSaveAddressesService struct {
	c             *Client
	saveAddressID int64
}

func (s *DeleteSaveAddressesService) SaveAddressID(saveAddressID int64) *DeleteSaveAddressesService {
	s.saveAddressID = saveAddressID
	return s
}

type DeleteAddressesResponse struct {
	basicReponse
	Result *string `json:"result"`
}

func (s *DeleteSaveAddressesService) Do(ctx context.Context) (*string, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/wallet/saved_addresses/%d", s.saveAddressID), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result DeleteAddressesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
