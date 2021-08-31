package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetLeveragedTokenInfoService struct {
	c     *Client
	token string
}

func (s *GetLeveragedTokenInfoService) Token(token string) *GetLeveragedTokenInfoService {
	s.token = token
	return s
}

type GetLeveragedTokenInfoResponse struct {
	basicReponse
	Result *LeveragedToken `json:"result"`
}

func (s *GetLeveragedTokenInfoService) Do(ctx context.Context) (*LeveragedToken, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/lt/%s", s.token), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetLeveragedTokenInfoResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
