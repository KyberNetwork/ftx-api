package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type CreateSaveAddressesService struct {
	c      *Client
	params CreateSaveAddressParams
}

type CreateSaveAddressParams struct {
	Coin         string  `json:"coin"`
	Address      string  `json:"address"`
	AddressName  string  `json:"addressName"`
	IsPrimetrust bool    `json:"isPrimetrust"`
	Tag          *string `json:"tag"`
}

func (s *CreateSaveAddressesService) Params(params CreateSaveAddressParams) *CreateSaveAddressesService {
	s.params = params
	return s
}

type CreateAddressesResponse struct {
	basicReponse
	Result *SaveAddress `json:"result"`
}

func (s *CreateSaveAddressesService) Do(ctx context.Context) (*SaveAddress, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/wallet/saved_addresses"), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result CreateAddressesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
