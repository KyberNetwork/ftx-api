package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetPositionsService struct {
	c            *Client
	showAvgPrice *bool
}

func (s *GetPositionsService) ShowAvgPrice(sap bool) *GetPositionsService {
	s.showAvgPrice = &sap
	return s
}

type PositionsResponse struct {
	basicReponse
	Result []Position `json:"result"`
}

func (s *GetPositionsService) Do(ctx context.Context) ([]Position, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/account"), true)
	if s.showAvgPrice != nil {
		r.setParam("showAvgPrice", BoolToString(*s.showAvgPrice))
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result PositionsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
