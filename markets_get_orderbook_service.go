package ftxapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type GetOrderBookService struct {
	c          *Client
	marketName string
	depth      *int
}

func (s *GetOrderBookService) MarketName(marketName string) *GetOrderBookService {
	s.marketName = marketName
	return s
}

func (s *GetOrderBookService) Depth(depth int) *GetOrderBookService {
	s.depth = &depth
	return s
}

type Feed struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

func (f *Feed) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&f.Price, &f.Size}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields: %d != %d", g, e)
	}
	return nil
}

type OrderBook struct {
	Asks []Feed `json:"asks"`
	Bids []Feed `json:"bids"`
}

type OrderBookResponse struct {
	basicReponse
	Result OrderBook `json:"result"`
}

func (s *GetOrderBookService) Do(ctx context.Context) (OrderBook, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/markets/%s/orderbook", s.marketName), false)
	if s.depth != nil {
		r.setParam("depth", IntToString(*s.depth))
	}
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return OrderBook{}, err
	}
	var result OrderBookResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return OrderBook{}, err
	}
	if !result.Success {
		return OrderBook{}, errors.New(result.Error)
	}
	return result.Result, nil
}
