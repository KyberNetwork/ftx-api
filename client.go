package ftxapi

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrorRateLimit = errors.New("error rate limit")
)

type Client struct {
	apiKey     string
	apiSecret  string
	baseURL    string
	httpClient *http.Client
	l          *zap.SugaredLogger
}

func NewClient(apiKey, apiSecret string, baseURL string, l *zap.SugaredLogger) *Client {
	return &Client{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		httpClient: http.DefaultClient,
		baseURL:    baseURL,
		l:          l,
	}
}

func (c *Client) callAPI(ctx context.Context, r *request) ([]byte, error) {
	req, err := c.parsedequest(ctx, r)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, errors.New("failed to read body")
	}
	if r.httpMethod != http.MethodGet {
		var rawData interface{}
		_ = json.Unmarshal(respBody, &rawData)
		c.l.Debugw("data response", "data", rawData)
	}
	if resp.StatusCode == 429 {
		return nil, ErrorRateLimit
	}
	if resp.StatusCode != http.StatusOK {
		var respData basicReponse
		if errU := json.Unmarshal(respBody, &respData); errU != nil {
			c.l.Errorw("cannot unmarshal response data", "err", err)
		} else {
			return nil, fmt.Errorf("unexpected status code = %d, error = %s", resp.StatusCode, respData.Error)
		}
		return nil, fmt.Errorf("unexpected status code = %d", resp.StatusCode)
	}
	return respBody, nil
}

func (c *Client) parsedequest(ctx context.Context, r *request) (*http.Request, error) {
	req, err := http.NewRequest(r.httpMethod, fmt.Sprintf("%s/%s", c.baseURL, r.endpoint), bytes.NewBuffer(r.body))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if r.httpMethod != http.MethodGet {
		req.Header.Set("Content-Type", "application/json")
	}
	query := req.URL.Query()
	for k, v := range r.params {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()
	if r.needSigned {
		nonce := fmt.Sprintf("%d", timeToTimestampMS(time.Now().UTC()))
		payload := nonce + req.Method + req.URL.Path
		if req.URL.RawQuery != "" {
			payload += "?" + req.URL.RawQuery
		}
		if len(r.body) > 0 {
			payload += string(r.body)
		}
		req.Header.Set("FTX-KEY", c.apiKey)
		req.Header.Set("FTX-TS", nonce)
		req.Header.Set("FTX-SIGN", c.sign(payload))
	}
	return req, nil
}

func (c *Client) sign(payload string) string {
	mac := hmac.New(sha256.New, []byte(c.apiSecret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

func (c *Client) NewGetSubAccountsService() *GetSubAccountsService {
	return &GetSubAccountsService{
		c: c,
	}
}

func (c *Client) NewGetMarketsService() *GetMarketsService {
	return &GetMarketsService{
		c: c,
	}
}

func (c *Client) NewGetSingleMarketsService() *GetSingleMarketService {
	return &GetSingleMarketService{
		c: c,
	}
}

func (c *Client) NewGetOrderBookService() *GetOrderBookService {
	return &GetOrderBookService{
		c: c,
	}
}

func (c *Client) NewGetTradesService() *GetTradesService {
	return &GetTradesService{
		c: c,
	}
}

func (c *Client) NewGetHistoricalPricesService() *GetHistoricalPricesService {
	return &GetHistoricalPricesService{
		c: c,
	}
}

func (c *Client) NewGetListFutureService() *GetListFutureService {
	return &GetListFutureService{
		c: c,
	}
}

func (c *Client) NewGetFutureService() *GetFutureService {
	return &GetFutureService{
		c: c,
	}
}

func (c *Client) NewGetFutureStatsService() *GetFutureStatsService {
	return &GetFutureStatsService{
		c: c,
	}
}

func (c *Client) NewGetFutureFundingRateService() *GetFutureFundingRateService {
	return &GetFutureFundingRateService{
		c: c,
	}
}

func (c *Client) NewGetFutureIndexWeightsService() *GetFutureIndexWeightsService {
	return &GetFutureIndexWeightsService{
		c: c,
	}
}

func (c *Client) NewGetExpiredFuturesService() *GetExpiredFuturesService {
	return &GetExpiredFuturesService{
		c: c,
	}
}

func (c *Client) NewGetHistoricalIndexService() *GetHistoricalIndexService {
	return &GetHistoricalIndexService{
		c: c,
	}
}

func (c *Client) NewGetAccountService() *GetAccountService {
	return &GetAccountService{
		c: c,
	}
}

func (c *Client) NewGetPositionsService() *GetPositionsService {
	return &GetPositionsService{
		c: c,
	}
}

func (c *Client) NewChangeAccountLeverageService() *ChangeAccountLeverageService {
	return &ChangeAccountLeverageService{
		c: c,
	}
}

//

func (c *Client) NewGetWithdrawalFeesService() *GetWithdrawalFeesService {
	return &GetWithdrawalFeesService{
		c: c,
	}
}
