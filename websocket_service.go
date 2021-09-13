package ftxapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	WebsocketEndpoint string = "wss://ftx.com/ws/"
)

type WebsocketService struct {
	l                     *zap.SugaredLogger
	mu                    sync.Mutex
	apiKey                string
	apiSecret             string
	wsEndpoint            string
	mapSubscriptions      map[Subscription]struct{}
	mapCheckSubscriptions map[Subscription]struct{}
	autoReconnect         bool
	conn                  *websocket.Conn
	subAccount            *string
	stopC                 chan struct{}
	stopPing              chan struct{}
	reconnectC            chan struct{}
	receivePong           chan struct{}
}

func NewWebsocketService(apiKey, apiSecret, wsEndpoint string, l *zap.SugaredLogger) *WebsocketService {
	return &WebsocketService{
		l:                l,
		mu:               sync.Mutex{},
		apiKey:           apiKey,
		apiSecret:        apiSecret,
		wsEndpoint:       wsEndpoint,
		mapSubscriptions: make(map[Subscription]struct{}),
		receivePong:      make(chan struct{}),
	}
}

func (s *WebsocketService) AutoReconnect() *WebsocketService {
	s.stopC = make(chan struct{})
	s.autoReconnect = true
	return s
}

func (s *WebsocketService) SubAccount(sa string) *WebsocketService {
	s.subAccount = &sa
	return s
}

func (s *WebsocketService) Connect(dataHandler WsDataHandler, errHandler WsErrorHandler) error {
	l := s.l.With("func", "WebsocketService.Connect")
	conn, _, err := websocket.DefaultDialer.Dial(s.wsEndpoint, nil)
	if err != nil {
		l.Errorw("cannot connect ws", "err", err)
		return err
	}

	// login
	if s.apiKey != "" && s.apiSecret != "" {
		if err := conn.WriteJSON(s.authenticationRequest()); err != nil {
			l.Errorw("failed to log in", "err", err)
			return err
		}
	}

	s.conn = conn
	s.mapCheckSubscriptions = make(map[Subscription]struct{})
	s.stopPing = make(chan struct{})
	s.reconnectC = make(chan struct{})

	go s.runPing()

	// subscribe
	for sub := range s.mapSubscriptions {
		if err := s.Subscribe(sub); err != nil {
			l.Errorw("failed to subscribe", "sub channel", sub.Channel,
				"market", sub.Market, "grouping", sub.Grouping, "err", err)
			return err
		}
	}

	go s.handleData(dataHandler, errHandler)
	if s.autoReconnect {
		go s.reconnect(dataHandler, errHandler)
	}
	return nil
}

func (s *WebsocketService) handleData(dataHandler WsDataHandler, errHandler WsErrorHandler) {
	defer func() {
		close(s.stopPing)
		close(s.reconnectC)
	}()
	l := s.l.With("func", "WebsocketService.handleData")
	for {
		_, msg, err := s.conn.ReadMessage()
		if err != nil {
			l.Errorw("cannot read msg from ws client", "err", err)
			errHandler(fmt.Errorf("cannot read msg from ws client, err = %s", err))
			return
		}
		j, err := simplejson.NewJson(msg)
		if err != nil {
			l.Errorw("cannot read json", "err", err)
			errHandler(fmt.Errorf("cannot read json, err = %s", err))
			return
		}

		typeEvent := j.Get("type").MustString("")
		channelEvent := j.Get("channel").MustString("")
		marketEvent := j.Get("market").MustString("")
		switch typeEvent {
		case "pong":
			s.receivePong <- struct{}{}
		case "subscribed":
			l.Infow("subscribe successfully", "channel", channelEvent, "market", marketEvent)
		case "unsubscribed":
			l.Infow("unsubscribe successfully", "channel", channelEvent, "market", marketEvent)
		case "error":
			var er errorResponse
			if err := json.Unmarshal(msg, &er); err != nil {
				l.Errorw("cannot unmarshal error data", "err", err)
				errHandler(fmt.Errorf("cannot unmarshal error data, err = %s", err))
				return
			}
			errHandler(fmt.Errorf("error from server, code = %d, msg = %s", er.Code, er.Msg))
		case "info":
			var info errorResponse
			if err := json.Unmarshal(msg, &info); err != nil {
				l.Errorw("cannot unmarshal info data", "err", err)
				errHandler(fmt.Errorf("cannot unmarshal info data, err = %s", err))
				return
			}
			if info.Code == 20001 {
				l.Infow("server suggest restart connection", "msg", info.Msg)
				return
			}
		case "partial", "update":
			switch WsChannel(channelEvent) {
			case WsChannelTicker:
				var event WsTickerEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					l.Errorw("cannot unmarshal orderbook data", "err", err)
					errHandler(fmt.Errorf("cannot unmarshal orderbook data, err = %s", err))
				}
				dataHandler(WsReponse{
					Ticker: &event,
				})
			case WsChannelMarkets:
				var event WsMarketsEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					l.Errorw("cannot unmarshal orderbook data", "err", err)
					errHandler(fmt.Errorf("cannot unmarshal orderbook data, err = %s", err))
				}
				dataHandler(WsReponse{
					Markets: &event,
				})
			case WsChannelTrades:
				var event WsTradesEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					l.Errorw("cannot unmarshal orderbook data", "err", err)
					errHandler(fmt.Errorf("cannot unmarshal orderbook data, err = %s", err))
				}
				dataHandler(WsReponse{
					Trades: &event,
				})
			case WsChannelOrderBook:
				var event WsOrderBookEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					l.Errorw("cannot unmarshal orderbook data", "err", err)
					errHandler(fmt.Errorf("cannot unmarshal orderbook data, err = %s", err))
				}
				dataHandler(WsReponse{
					OrderBookEvent: &event,
				})
			case WsChannelOrderbookGrouped:
				var event WsGroupedOrderBookEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					l.Errorw("cannot unmarshal grouped orderbook data", "err", err)
					errHandler(fmt.Errorf("cannot unmarshal grouped orderbook data, err = %s", err))
				}
				dataHandler(WsReponse{
					GroupedOrderBookEvent: &event,
				})
			case WsChannelFills:
				var event WsFillsEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					l.Errorw("cannot unmarshal fills data", "err", err)
					errHandler(fmt.Errorf("cannot unmarshal fills data, err = %s", err))
				}
				dataHandler(WsReponse{
					Fills: &event,
				})
			case WsChannelOrders:
				var event WsOrdersEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					l.Errorw("cannot unmarshal orders data", "err", err)
					errHandler(fmt.Errorf("cannot unmarshal orders data, err = %s", err))
				}
				dataHandler(WsReponse{
					Orders: &event,
				})
			case WsChannelFTXPay:
				var event WsFTXPayEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					l.Errorw("cannot unmarshal ftx pay data", "err", err)
					errHandler(fmt.Errorf("cannot unmarshal ftx pay data, err = %s", err))
				}
				dataHandler(WsReponse{
					FTXPay: &event,
				})
			}
		default:
			l.Infow("not supported type", "type", typeEvent)
		}
	}
}

func (s *WebsocketService) authenticationRequest() RequestMsg {
	t := timeToTimestampMS(time.Now())
	mac := hmac.New(sha256.New, []byte(s.apiSecret))
	mac.Write([]byte(fmt.Sprintf("%dwebsocket_login", t)))
	signature := hex.EncodeToString(mac.Sum(nil))
	args := map[string]interface{}{
		"key":  s.apiKey,
		"time": t,
		"sign": signature,
	}
	if s.subAccount != nil {
		args["subaccount"] = *s.subAccount
	}
	return RequestMsg{
		OP:   "login",
		Args: args,
	}
}

func (s *WebsocketService) closeConnection() {
	_ = s.conn.Close()
}

func (s *WebsocketService) runPing() {
	t := time.NewTicker(15 * time.Second)
	defer t.Stop()
loop:
	for {
		select {
		case <-t.C:
			tm := time.NewTimer(2 * time.Second)
			if err := s.conn.WriteJSON(RequestMsg{
				OP: "ping",
			}); err != nil {
				s.closeConnection()
				break loop
			}
			select {
			case <-s.receivePong:
				tm.Stop()
			case <-tm.C:
				s.closeConnection()
				break loop
			}
		case <-s.stopPing:
			break loop
		}
	}
}

func (s *WebsocketService) Subscribe(sub Subscription) error {
	s.mu.Lock()
	mapCheckSub := s.mapCheckSubscriptions
	s.mu.Unlock()
	if _, ok := mapCheckSub[sub]; ok {
		return nil
	}
	if err := s.conn.WriteJSON(RequestMsg{
		OP:       "subscribe",
		Channel:  StringPointer(string(sub.Channel)),
		Market:   sub.Market,
		Grouping: sub.Grouping,
	}); err != nil {
		return err
	}
	s.mu.Lock()
	s.mapCheckSubscriptions[sub] = struct{}{}
	s.mapSubscriptions[sub] = struct{}{}
	s.mu.Unlock()
	return nil
}

func (s *WebsocketService) Unsubscribe(sub Subscription) error {
	s.mu.Lock()
	mapCheckSub := s.mapCheckSubscriptions
	s.mu.Unlock()
	if _, ok := mapCheckSub[sub]; !ok {
		return nil
	}
	if err := s.conn.WriteJSON(RequestMsg{
		OP:       "unsubscribe",
		Channel:  StringPointer(string(sub.Channel)),
		Market:   sub.Market,
		Grouping: sub.Grouping,
	}); err != nil {
		return err
	}
	s.mu.Lock()
	delete(s.mapCheckSubscriptions, sub)
	delete(s.mapSubscriptions, sub)
	s.mu.Unlock()
	return nil
}

func (s *WebsocketService) reconnect(dataHandler WsDataHandler, errHandler WsErrorHandler) {
	l := s.l.With("func", "reconnect")
	select {
	case <-s.stopC:
		l.Infow("connection will be stopped")
	case <-s.reconnectC:
		l.Infow("reconnect...")
		time.Sleep(1 * time.Second)
		for {
			if err := s.Connect(dataHandler, errHandler); err != nil {
				l.Errorw("reconnect: connect error", "err", err)
				time.Sleep(5 * time.Second)
			} else {
				l.Infow("reconnect successfully")
				break
			}
		}
	}
}

func (s *WebsocketService) ResetConnection() {
	s.closeConnection()
}

func (s *WebsocketService) Close() {
	if s.autoReconnect {
		close(s.stopC)
	}
	s.closeConnection()
}
