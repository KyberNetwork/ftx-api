package ftxapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/websocket"
)

const (
	websocketEndpoint string = "wss://ftx.com/ws/"
)

type WebsocketService struct {
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

func NewWebsocketService(apiKey, apiSecret string, autoReconnect bool) *WebsocketService {
	s := &WebsocketService{
		mu:               sync.Mutex{},
		apiKey:           apiKey,
		apiSecret:        apiSecret,
		wsEndpoint:       websocketEndpoint,
		mapSubscriptions: make(map[Subscription]struct{}),
		receivePong:      make(chan struct{}),
		autoReconnect:    autoReconnect,
	}
	if autoReconnect {
		s.stopC = make(chan struct{})
	}
	return s
}

func (s *WebsocketService) SubAccount(sa string) *WebsocketService {
	s.subAccount = &sa
	return s
}

func (s *WebsocketService) Connect(dataHandler WsDataHandler, errHandler WsErrorHandler) error {
	conn, _, err := websocket.DefaultDialer.Dial(s.wsEndpoint, nil)
	if err != nil {
		log.Printf("failed to connect ws, err = %s\n", err)
		return err
	}

	// login
	if s.apiKey != "" && s.apiSecret != "" {
		if err := conn.WriteJSON(s.authenticationRequest()); err != nil {
			log.Printf("failed to write authentication request, err = %s\n", err)
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
			log.Printf("failed to subscribe, channel = %s, market = %v, grouping = %v, err = %s\n", sub.Channel, sub.Market, sub.Grouping, err)
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
	for {
		_, msg, err := s.conn.ReadMessage()
		if err != nil {
			errHandler(fmt.Errorf("failed to read msg from ws client, err = %s", err))
			return
		}
		j, err := simplejson.NewJson(msg)
		if err != nil {
			errHandler(fmt.Errorf("failed to read json, err = %s", err))
			return
		}

		typeEvent := j.Get("type").MustString("")
		channelEvent := j.Get("channel").MustString("")
		marketEvent := j.Get("market").MustString("")
		switch typeEvent {
		case "pong":
			s.receivePong <- struct{}{}
		case "subscribed":
			log.Printf("subscribed successfully, channel event = %s, market = %s\n", channelEvent, marketEvent)
		case "unsubscribed":
			log.Printf("unsubscribed successfully, channel event = %s, market = %s\n", channelEvent, marketEvent)
		case "error":
			var er errorResponse
			if err := json.Unmarshal(msg, &er); err != nil {
				errHandler(fmt.Errorf("failed to unmarshal error data, err = %s", err))
				return
			}
			errHandler(fmt.Errorf("error from server, code = %d, msg = %s", er.Code, er.Msg))
		case "info":
			var info errorResponse
			if err := json.Unmarshal(msg, &info); err != nil {
				errHandler(fmt.Errorf("failed to unmarshal info data, err = %s", err))
				return
			}
			if info.Code == 20001 {
				log.Printf("server suggests to restart the connection msg = %s\n", info.Msg)
				return
			}
		case "partial", "update":
			switch WsChannel(channelEvent) {
			case WsChannelTicker:
				var event WsTickerEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					errHandler(fmt.Errorf("failed to unmarshal ticker event, err = %s", err))
				} else {
					dataHandler(WsReponse{
						Ticker: &event,
					})
				}
			case WsChannelMarkets:
				var event WsMarketsEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					errHandler(fmt.Errorf("failed to unmarshal market event, err = %s", err))
				} else {
					dataHandler(WsReponse{
						Markets: &event,
					})
				}
			case WsChannelTrades:
				var event WsTradesEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					errHandler(fmt.Errorf("failed to unmarshal trade event, err = %s", err))
				} else {
					dataHandler(WsReponse{
						Trades: &event,
					})
				}
			case WsChannelOrderBook:
				var event WsOrderBookEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					errHandler(fmt.Errorf("failed to unmarshal orderbook event, err = %s", err))
				} else {
					dataHandler(WsReponse{
						OrderBookEvent: &event,
					})
				}
			case WsChannelOrderbookGrouped:
				var event WsGroupedOrderBookEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					errHandler(fmt.Errorf("failed to unmarshal grouped orderbook event, err = %s", err))
				} else {
					dataHandler(WsReponse{
						GroupedOrderBookEvent: &event,
					})
				}
			case WsChannelFills:
				var event WsFillsEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					errHandler(fmt.Errorf("failed to unmarshal fills event, err = %s", err))
				} else {
					dataHandler(WsReponse{
						Fills: &event,
					})
				}
			case WsChannelOrders:
				var event WsOrdersEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					errHandler(fmt.Errorf("failed to unmarshal orders event, err = %s", err))
				} else {
					dataHandler(WsReponse{
						Orders: &event,
					})
				}
			case WsChannelFTXPay:
				var event WsFTXPayEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					errHandler(fmt.Errorf("failed to unmarshal ftx pay event, err = %s", err))
				} else {
					dataHandler(WsReponse{
						FTXPay: &event,
					})
				}
			}
		default:
			log.Printf("event type is not supported, event = %s\n", typeEvent)
		}
	}
}

func (s *WebsocketService) authenticationRequest() RequestMsg {
	t := time.Now().UnixMilli()
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

func (s *WebsocketService) isInCheckSub(sub Subscription) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.mapCheckSubscriptions[sub]
	return ok
}

func (s *WebsocketService) Subscribe(sub Subscription) error {
	if s.isInCheckSub(sub) {
		return nil
	}
	if err := s.conn.WriteJSON(RequestMsg{
		OP:       "subscribe",
		Channel:  StringToPointer(string(sub.Channel)),
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
	if s.isInCheckSub(sub) {
		return nil
	}
	if err := s.conn.WriteJSON(RequestMsg{
		OP:       "unsubscribe",
		Channel:  StringToPointer(string(sub.Channel)),
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
	select {
	case <-s.stopC:
		log.Printf("websocket is closed\n")
	case <-s.reconnectC:
		log.Printf("reconnecting...\n")
		time.Sleep(1 * time.Second)
		for {
			if err := s.Connect(dataHandler, errHandler); err != nil {
				log.Printf("failed to reconnect err = %s\n", err)
				time.Sleep(5 * time.Second)
			} else {
				log.Printf("reconnected successfully\n")
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
