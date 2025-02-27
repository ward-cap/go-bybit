package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// V5WebsocketPublicServiceI :
type V5WebsocketPublicServiceI interface {
	Start(context.Context) error
	//Run() error
	Close() error

	SubscribeOrderBook(
		V5WebsocketPublicOrderBookParamKey,
		func(V5WebsocketPublicOrderBookResponse) error,
	) (func() error, error)

	SubscribeKline(
		V5WebsocketPublicKlineParamKey,
		func(V5WebsocketPublicKlineResponse) error,
	) (func() error, error)

	SubscribeTicker(
		V5WebsocketPublicTickerParamKey,
		func(V5WebsocketPublicTickerResponse) error,
	) (func() error, error)

	SubscribeTrade(
		V5WebsocketPublicTradeParamKey,
		func(V5WebsocketPublicTradeResponse) error,
	) (func() error, error)

	SubscribeLiquidation(
		V5WebsocketPublicLiquidationParamKey,
		func(V5WebsocketPublicLiquidationResponse) error,
	) (func() error, error)
}

// V5WebsocketPublicService :
type V5WebsocketPublicService struct {
	client     *WebSocketClient
	connection *websocket.Conn
	category   CategoryV5

	mu sync.Mutex

	logger *zap.SugaredLogger

	paramOrderBookMap   map[V5WebsocketPublicOrderBookParamKey]func(V5WebsocketPublicOrderBookResponse) error
	paramKlineMap       map[V5WebsocketPublicKlineParamKey]func(V5WebsocketPublicKlineResponse) error
	paramTickerMap      map[V5WebsocketPublicTickerParamKey]func(V5WebsocketPublicTickerResponse) error
	paramTradeMap       map[V5WebsocketPublicTradeParamKey]func(V5WebsocketPublicTradeResponse) error
	paramLiquidationMap map[V5WebsocketPublicLiquidationParamKey]func(V5WebsocketPublicLiquidationResponse) error
}

const (
	// V5WebsocketPublicPath :
	V5WebsocketPublicPath = "/v5/public"
)

// V5WebsocketPublicPathFor :
func V5WebsocketPublicPathFor(category CategoryV5) string {
	return V5WebsocketPublicPath + "/" + string(category)
}

// V5WebsocketPublicTopic :
type V5WebsocketPublicTopic string

const (
	// V5WebsocketPublicTopicOrderBook :
	V5WebsocketPublicTopicOrderBook = V5WebsocketPublicTopic("orderbook")

	// V5WebsocketPublicTopicKline :
	V5WebsocketPublicTopicKline = V5WebsocketPublicTopic("kline")

	// V5WebsocketPublicTopicTicker :
	V5WebsocketPublicTopicTicker = V5WebsocketPublicTopic("tickers")

	// V5WebsocketPublicTopicTrade :
	V5WebsocketPublicTopicTrade = V5WebsocketPublicTopic("publicTrade")

	// V5WebsocketPublicTopicLiquidation :
	V5WebsocketPublicTopicLiquidation = V5WebsocketPublicTopic("liquidation")
)

func (t V5WebsocketPublicTopic) String() string {
	return string(t)
}

// judgeTopic :
func (s *V5WebsocketPublicService) judgeTopic(respBody []byte) (V5WebsocketPublicTopic, error) {
	parsedData := map[string]any{}
	if err := json.Unmarshal(respBody, &parsedData); err != nil {
		return "", err
	}
	if topic, ok := parsedData["topic"].(string); ok {
		switch {
		case strings.Contains(topic, V5WebsocketPublicTopicOrderBook.String()):
			return V5WebsocketPublicTopicOrderBook, nil
		case strings.Contains(topic, V5WebsocketPublicTopicKline.String()):
			return V5WebsocketPublicTopicKline, nil
		case strings.Contains(topic, V5WebsocketPublicTopicTicker.String()):
			return V5WebsocketPublicTopicTicker, nil
		case strings.Contains(topic, V5WebsocketPublicTopicTrade.String()):
			return V5WebsocketPublicTopicTrade, nil
		case strings.Contains(topic, V5WebsocketPublicTopicLiquidation.String()):
			return V5WebsocketPublicTopicLiquidation, nil
		default:
			s.logger.Warnf("Unhandled topic: %s", topic)
		}
	}
	return "", nil
}

// UnmarshalJSON :
func (r *V5WebsocketPublicTickerData) UnmarshalJSON(data []byte) error {
	switch r.category {
	case CategoryV5Linear, CategoryV5Inverse:
		return json.Unmarshal(data, &r.LinearInverse)
	case CategoryV5Option:
		return json.Unmarshal(data, &r.Option)
	case CategoryV5Spot:
		return json.Unmarshal(data, &r.Spot)
	}
	return errors.New("unsupported format")
}

// parseResponse :
func (s *V5WebsocketPublicService) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// Start :
func (s *V5WebsocketPublicService) Start(ctx context.Context) error {

	go func() {
		defer s.Close()

		s.keepAlive(s.connection)

		for {
			if err := s.Run(); err != nil {
				break
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			s.logger.Debug("caught websocket public service interrupt signal")
			return s.Close()
		}
	}
}

// Run :
func (s *V5WebsocketPublicService) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	topic, err := s.judgeTopic(message)
	if err != nil {
		return err
	}
	switch topic {
	//case V5WebsocketPrivateTopicPong:

	case V5WebsocketPublicTopicOrderBook:
		var resp V5WebsocketPublicOrderBookResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrieveOrderBookFunc(resp.Key())
		if err != nil {
			return err
		}
		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPublicTopicKline:
		var resp V5WebsocketPublicKlineResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		f, err := s.retrieveKlineFunc(resp.Key())
		if err != nil {
			return err
		}

		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPublicTopicTicker:
		var resp V5WebsocketPublicTickerResponse
		resp.Data.category = s.category
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		f, err := s.retrieveTickerFunc(resp.Key())
		if err != nil {
			return err
		}

		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPublicTopicTrade:
		var resp V5WebsocketPublicTradeResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		f, err := s.retrieveTradeFunc(resp.Key())
		if err != nil {
			return err
		}

		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPublicTopicLiquidation:
		var resp V5WebsocketPublicLiquidationResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		f, err := s.retrieveLiquidationFunc(resp.Key())
		if err != nil {
			return err
		}

		if err := f(resp); err != nil {
			return err
		}
	}
	return nil
}

// Close :
func (s *V5WebsocketPublicService) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.connection.Close()
}

func (s *V5WebsocketPublicService) writeMessage(messageType int, body []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.connection.WriteMessage(messageType, body)
}

func (s *V5WebsocketPublicService) keepAlive(c *websocket.Conn) {
	timeout := time.Second * 10
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		s.logger.Debug("pong")
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			go func() {
				err := errors.Join(
					s.writeMessage(websocket.PingMessage, nil),
					s.writeMessage(websocket.TextMessage, []byte(`{"op":"ping"}`)),
				)

				if err != nil {
					_ = c.Close()
					return
				}
			}()

			<-ticker.C
			if time.Since(lastResponse) > timeout {
				_ = c.Close()
				return
			}
		}
	}()
}
