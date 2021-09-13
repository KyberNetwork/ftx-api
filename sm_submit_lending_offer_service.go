package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type SubmitLendingOfferService struct {
	c      *Client
	params SubmitLendingOfferParams
}

type SubmitLendingOfferParams struct {
	Coin string  `json:"coin"`
	Size float64 `json:"size"`
	Rate float64 `json:"rate"`
}

func (s *SubmitLendingOfferService) Params(params SubmitLendingOfferParams) *SubmitLendingOfferService {
	s.params = params
	return s
}

type SubmitLendingOfferResponse struct {
	basicReponse
	Result interface{} `json:"result"`
}

func (s *SubmitLendingOfferService) Do(ctx context.Context) (interface{}, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/spot_margin/offers"), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result SubmitLendingOfferResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
