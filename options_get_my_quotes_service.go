package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetMyQuotesService struct {
	c *Client
}

type GetMyQuotesResponse struct {
	basicReponse
	Result []OptionQuote `json:"result"`
}

func (s *GetMyQuotesService) Do(ctx context.Context) ([]OptionQuote, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/options/my_quotes"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result GetMyQuotesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
