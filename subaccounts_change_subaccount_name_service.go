package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type ChangeSubAccountNameService struct {
	c      *Client
	params ChangeSubAccountNameParams
}

type ChangeSubAccountNameParams struct {
	Nickname    string `json:"nickname"`
	NewNickName string `json:"newNickname"`
}

func (s *ChangeSubAccountNameService) Params(params ChangeSubAccountNameParams) *ChangeSubAccountNameService {
	s.params = params
	return s
}

type ChangeSubAccountNameResponse struct {
	basicReponse
}

func (s *ChangeSubAccountNameService) Do(ctx context.Context) error {
	r := newRequest(http.MethodPost, "/subaccounts/update_name", true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return err
	}
	var result ChangeSubAccountNameResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return err
	}
	if !result.Success {
		return errors.New(result.Error)
	}
	return nil
}
