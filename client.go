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
	ErrorRateLimit  = errors.New("error rate limit")
	RestAPIEndpoint = "https://ftx.com/api"
)

type Client struct {
	l          *zap.SugaredLogger
	apiKey     string
	apiSecret  string
	baseURL    string
	httpClient *http.Client
	subAccount *string
}

func NewClient(apiKey, apiSecret, baseURL string, l *zap.SugaredLogger) *Client {
	return &Client{
		l:          l,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		httpClient: http.DefaultClient,
		baseURL:    baseURL,
	}
}

func (c *Client) SubAccount(subaccount *string) *Client {
	c.subAccount = subaccount
	return c
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
		nonce := fmt.Sprintf("%d", time.Now().UnixMilli())
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
	if c.subAccount != nil {
		req.Header.Set("FTX-SUBACCOUNT", *c.subAccount)
	}
	return req, nil
}

func (c *Client) sign(payload string) string {
	mac := hmac.New(sha256.New, []byte(c.apiSecret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

func (c *Client) NewGetAllSubAccountsService() *GetAllSubAccountsService {
	return &GetAllSubAccountsService{
		c: c,
	}
}

func (c *Client) NewCreateSubAccountService() *CreateSubAccountService {
	return &CreateSubAccountService{
		c: c,
	}
}

func (c *Client) NewChangeSubAccountNameService() *ChangeSubAccountNameService {
	return &ChangeSubAccountNameService{
		c: c,
	}
}

func (c *Client) NewDeleteSubAccountService() *DeleteSubAccountService {
	return &DeleteSubAccountService{
		c: c,
	}
}

func (c *Client) NewTransferBetweenSubAccountsService() *TransferBetweenSubAccountsService {
	return &TransferBetweenSubAccountsService{
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

func (c *Client) NewGetCoinsService() *GetCoinsService {
	return &GetCoinsService{
		c: c,
	}
}

func (c *Client) NewGetBalancesService() *GetBalancesService {
	return &GetBalancesService{
		c: c,
	}
}

func (c *Client) NewGetAllBalancesService() *GetAllBalancesService {
	return &GetAllBalancesService{
		c: c,
	}
}

func (c *Client) NewGetDepositAddressService() *GetDepositAddressService {
	return &GetDepositAddressService{
		c: c,
	}
}

func (c *Client) NewGetDepositAddressListService() *GetDepositAddressListService {
	return &GetDepositAddressListService{
		c: c,
	}
}

func (c *Client) NewGetDepositHistoryService() *GetDepositHistoryService {
	return &GetDepositHistoryService{
		c: c,
	}
}

func (c *Client) NewGetWithdrawHistoryService() *GetWithdrawHistoryService {
	return &GetWithdrawHistoryService{
		c: c,
	}
}

func (c *Client) NewWithdrawService() *WithdrawService {
	return &WithdrawService{
		c: c,
	}
}

func (c *Client) NewGetAirdropsService() *GetAirdropsService {
	return &GetAirdropsService{
		c: c,
	}
}

func (c *Client) NewGetWithdrawalFeesService() *GetWithdrawalFeesService {
	return &GetWithdrawalFeesService{
		c: c,
	}
}

func (c *Client) NewGetSaveAddressesService() *GetSaveAddressesService {
	return &GetSaveAddressesService{
		c: c,
	}
}

func (c *Client) NewCreateSaveAddressesService() *CreateSaveAddressesService {
	return &CreateSaveAddressesService{
		c: c,
	}
}

func (c *Client) NewDeleteSaveAddressesService() *DeleteSaveAddressesService {
	return &DeleteSaveAddressesService{
		c: c,
	}
}

func (c *Client) NewGetOpenOrdersService() *GetOpenOrdersService {
	return &GetOpenOrdersService{
		c: c,
	}
}

func (c *Client) NewGetOrderHistoryService() *GetOrderHistoryService {
	return &GetOrderHistoryService{
		c: c,
	}
}

func (c *Client) NewGetOpenTriggerOrdersService() *GetOpenTriggerOrdersService {
	return &GetOpenTriggerOrdersService{
		c: c,
	}
}

func (c *Client) NewGetTriggerOrderHistoryService() *GetTriggerOrderHistoryService {
	return &GetTriggerOrderHistoryService{
		c: c,
	}
}

func (c *Client) NewPlaceOrderService() *PlaceOrderService {
	return &PlaceOrderService{
		c: c,
	}
}

func (c *Client) NewPlaceTriggerOrderService() *PlaceTriggerOrderService {
	return &PlaceTriggerOrderService{
		c: c,
	}
}

func (c *Client) NewModifyOrderService() *ModifyOrderService {
	return &ModifyOrderService{
		c: c,
	}
}

func (c *Client) NewModifyOrderByClientIDService() *ModifyOrderByClientIDService {
	return &ModifyOrderByClientIDService{
		c: c,
	}
}

func (c *Client) NewModifyTriggerOrderService() *ModifyTriggerOrderService {
	return &ModifyTriggerOrderService{
		c: c,
	}
}

func (c *Client) NewGetOrderStatusService() *GetOrderStatusService {
	return &GetOrderStatusService{
		c: c,
	}
}

func (c *Client) NewGetOrderStatusByClientIDService() *GetOrderStatusByClientIDService {
	return &GetOrderStatusByClientIDService{
		c: c,
	}
}

func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{
		c: c,
	}
}

func (c *Client) NewCancelOrderByClientIDService() *CancelOrderByClientIDService {
	return &CancelOrderByClientIDService{
		c: c,
	}
}

func (c *Client) NewCancelTriggerOrderService() *CancelTriggerOrderService {
	return &CancelTriggerOrderService{
		c: c,
	}
}

func (c *Client) NewCancelAllOrderService() *CancelAllOrderService {
	return &CancelAllOrderService{
		c: c,
	}
}

func (c *Client) NewFundingPaymentsService() *FundingPaymentsService {
	return &FundingPaymentsService{
		c: c,
	}
}

func (c *Client) NewListLeveragedTokensService() *ListLeveragedTokensService {
	return &ListLeveragedTokensService{
		c: c,
	}
}

func (c *Client) NewGetLeveragedTokenInfoService() *GetLeveragedTokenInfoService {
	return &GetLeveragedTokenInfoService{
		c: c,
	}
}

func (c *Client) NewGetLeveragedTokenBalancesService() *GetLeveragedTokenBalancesService {
	return &GetLeveragedTokenBalancesService{
		c: c,
	}
}

func (c *Client) NewListLeveragedTokenCreationRequestsService() *ListLeveragedTokenCreationRequestsService {
	return &ListLeveragedTokenCreationRequestsService{
		c: c,
	}
}

func (c *Client) NewRequestLeveragedTokenCreationService() *RequestLeveragedTokenCreationService {
	return &RequestLeveragedTokenCreationService{
		c: c,
	}
}

func (c *Client) NewListLeveragedTokenRedemptionRequestsService() *ListLeveragedTokenRedemptionRequestsService {
	return &ListLeveragedTokenRedemptionRequestsService{
		c: c,
	}
}

func (c *Client) NewRequestLeveragedTokenRedemptionService() *RequestLeveragedTokenRedemptionService {
	return &RequestLeveragedTokenRedemptionService{
		c: c,
	}
}

func (c *Client) NewRequestETFRebalanceInfoService() *RequestETFRebalanceInfoService {
	return &RequestETFRebalanceInfoService{
		c: c,
	}
}

func (c *Client) NewListQuoteRequestsService() *ListQuoteRequestsService {
	return &ListQuoteRequestsService{
		c: c,
	}
}

func (c *Client) NewYourQuoteRequestsService() *YourQuoteRequestsService {
	return &YourQuoteRequestsService{
		c: c,
	}
}

func (c *Client) NewCreateQuoteRequestService() *CreateQuoteRequestService {
	return &CreateQuoteRequestService{
		c: c,
	}
}

func (c *Client) NewCancelQuoteRequestService() *CancelQuoteRequestService {
	return &CancelQuoteRequestService{
		c: c,
	}
}

func (c *Client) NewGetQuotesForYourQuoteRequestService() *GetQuotesForYourQuoteRequestService {
	return &GetQuotesForYourQuoteRequestService{
		c: c,
	}
}

func (c *Client) NewCreateQuoteService() *CreateQuoteService {
	return &CreateQuoteService{
		c: c,
	}
}

func (c *Client) NewGetMyQuotesService() *GetMyQuotesService {
	return &GetMyQuotesService{
		c: c,
	}
}

func (c *Client) NewCancelQuoteService() *CancelQuoteService {
	return &CancelQuoteService{
		c: c,
	}
}

func (c *Client) NewAcceptOptionsQuoteService() *AcceptOptionsQuoteService {
	return &AcceptOptionsQuoteService{
		c: c,
	}
}

func (c *Client) NewGetAccountOptionsInfoService() *GetAccountOptionsInfoService {
	return &GetAccountOptionsInfoService{
		c: c,
	}
}

func (c *Client) NewGetPublicOptionsTradesService() *GetPublicOptionsTradesService {
	return &GetPublicOptionsTradesService{
		c: c,
	}
}

func (c *Client) NewGet24HOptionVolumeService() *Get24HOptionVolumeService {
	return &Get24HOptionVolumeService{
		c: c,
	}
}

func (c *Client) NewGetHistorical24HOptionVolumeService() *GetHistorical24HOptionVolumeService {
	return &GetHistorical24HOptionVolumeService{
		c: c,
	}
}

func (c *Client) NewGetOptionOpenInterestService() *GetOptionOpenInterestService {
	return &GetOptionOpenInterestService{
		c: c,
	}
}

func (c *Client) NewGetHistoricalOpenInterestService() *GetHistoricalOpenInterestService {
	return &GetHistoricalOpenInterestService{
		c: c,
	}
}

func (c *Client) NewGetLendingHistoryService() *GetLendingHistoryService {
	return &GetLendingHistoryService{
		c: c,
	}
}

func (c *Client) NewGetBorrowRatesService() *GetBorrowRatesService {
	return &GetBorrowRatesService{
		c: c,
	}
}

func (c *Client) NewGetLendingRatesService() *GetLendingRatesService {
	return &GetLendingRatesService{
		c: c,
	}
}

func (c *Client) NewGetDailyBorrowedAmountsService() *GetDailyBorrowedAmountsService {
	return &GetDailyBorrowedAmountsService{
		c: c,
	}
}

func (c *Client) NewGetSpotMarginMarketInfoService() *GetSpotMarginMarketInfoService {
	return &GetSpotMarginMarketInfoService{
		c: c,
	}
}

func (c *Client) NewGetMyBorrowHistoryService() *GetMyBorrowHistoryService {
	return &GetMyBorrowHistoryService{
		c: c,
	}
}

func (c *Client) NewGetMyLendingHistoryService() *GetMyLendingHistoryService {
	return &GetMyLendingHistoryService{
		c: c,
	}
}

func (c *Client) NewGetLendingOffersService() *GetLendingOffersService {
	return &GetLendingOffersService{
		c: c,
	}
}

func (c *Client) NewGetLendingInfoService() *GetLendingInfoService {
	return &GetLendingInfoService{
		c: c,
	}
}

func (c *Client) NewSubmitLendingOfferService() *SubmitLendingOfferService {
	return &SubmitLendingOfferService{
		c: c,
	}
}
