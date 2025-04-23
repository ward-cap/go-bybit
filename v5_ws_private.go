package bybit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// V5WebsocketPrivateServiceI :
type V5WebsocketPrivateServiceI interface {
	Start(context.Context) error
	Subscribe() error
	Close() error

	SubscribeOrder(
		func(V5WebsocketPrivateOrderResponse) error,
	) (func() error, error)

	SubscribePosition(
		func(V5WebsocketPrivatePositionResponse) error,
	) (func() error, error)

	SubscribeExecution(
		func(V5WebsocketPrivateExecutionResponse) error,
	) (func() error, error)

	SubscribeWallet(
		func(V5WebsocketPrivateWalletResponse) error,
	) (func() error, error)
}

// V5WebsocketPrivateService :
type V5WebsocketPrivateService struct {
	client     *WebSocketClient
	connection *websocket.Conn

	logger *zap.SugaredLogger

	mu sync.Mutex

	paramOrderMap     map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateOrderResponse) error
	paramPositionMap  map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivatePositionResponse) error
	paramExecutionMap map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateExecutionResponse) error
	paramWalletMap    map[V5WebsocketPrivateParamKey]func(V5WebsocketPrivateWalletResponse) error
}

const (
	// V5WebsocketPrivatePath :
	V5WebsocketPrivatePath = "/v5/private"
)

// V5WebsocketPrivateTopic :
type V5WebsocketPrivateTopic string

const (
	// V5WebsocketPrivateTopicPong :
	V5WebsocketPrivateTopicPong V5WebsocketPrivateTopic = "pong"

	// V5WebsocketPrivateTopicOrder :
	V5WebsocketPrivateTopicOrder V5WebsocketPrivateTopic = "order"

	// V5WebsocketPrivateTopicPosition :
	V5WebsocketPrivateTopicPosition V5WebsocketPrivateTopic = "position"

	// V5WebsocketPrivateTopicExecution :
	V5WebsocketPrivateTopicExecution V5WebsocketPrivateTopic = "execution"

	// V5WebsocketPrivateTopicWallet :
	V5WebsocketPrivateTopicWallet V5WebsocketPrivateTopic = "wallet"
)

// V5WebsocketPrivateParamKey :
type V5WebsocketPrivateParamKey struct {
	Topic V5WebsocketPrivateTopic
}

// judgeTopic :
func (s *V5WebsocketPrivateService) judgeTopic(respBody []byte) (V5WebsocketPrivateTopic, error) {
	parsedData := map[string]interface{}{}
	if err := json.Unmarshal(respBody, &parsedData); err != nil {
		return "", err
	}
	if retMsg, ok := parsedData["op"].(string); ok && retMsg == "pong" {
		return V5WebsocketPrivateTopicPong, nil
	}
	if topic, ok := parsedData["topic"].(string); ok {
		return V5WebsocketPrivateTopic(topic), nil
	}
	if authStatus, ok := parsedData["success"].(bool); ok {
		if !authStatus {
			return "", errors.New("auth failed: " + parsedData["ret_msg"].(string))
		}
	}
	return "", nil
}

// parseResponse :
func (s *V5WebsocketPrivateService) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// Subscribe : Apply for authentication when establishing a connection.
func (s *V5WebsocketPrivateService) Subscribe() error {
	param, err := s.client.buildAuthParam()
	if err != nil {
		return err
	}
	if err := s.writeMessage(websocket.TextMessage, param); err != nil {
		return err
	}
	return nil
}

// ErrHandler :
type ErrHandler func(isWebsocketClosed bool, err error)

// Start :
func (s *V5WebsocketPrivateService) Start(ctx context.Context) error {
	ctx, cancelFunc := context.WithCancel(ctx)
	go func() {
		defer s.Close()

		s.keepAlive(s.connection, cancelFunc)

		for {
			if err := s.Run(); err != nil {
				s.logger.Error(err)
				return
			}
		}
	}()

	select {
	case <-ctx.Done():
		s.logger.Debug("caught websocket private service interrupt signal")
		return s.Close()
	}
}

// Run :
func (s *V5WebsocketPrivateService) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	topic, err := s.judgeTopic(message)
	if err != nil {
		return err
	}
	switch topic {
	case V5WebsocketPrivateTopicPong:
		if err := s.connection.PongHandler()("pong"); err != nil {
			return fmt.Errorf("pong: %w", err)
		}
	case V5WebsocketPrivateTopicOrder:
		var resp V5WebsocketPrivateOrderResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrieveOrderFunc(resp.Key())
		if err != nil {
			return err
		}
		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPrivateTopicPosition:
		var resp V5WebsocketPrivatePositionResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrievePositionFunc(resp.Key())
		if err != nil {
			return err
		}
		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPrivateTopicExecution:
		var resp V5WebsocketPrivateExecutionResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrieveExecutionFunc(resp.Key())
		if err != nil {
			return err
		}
		if err := f(resp); err != nil {
			return err
		}
	case V5WebsocketPrivateTopicWallet:
		var resp V5WebsocketPrivateWalletResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		f, err := s.retrieveWalletFunc(resp.Key())
		if err != nil {
			return err
		}
		if err := f(resp); err != nil {
			return err
		}
	}

	return nil
}

// Ping :
//func (s *V5WebsocketPrivateService) Ping() error {
//	// NOTE: It appears that two messages need to be sent.
//	// REF: https://github.com/hirokisan/bybit/pull/127#issuecomment-1537479346
//	if err := s.writeMessage(websocket.PingMessage, nil); err != nil {
//		return err
//	}
//	if err := s.writeMessage(websocket.TextMessage, []byte(`{"op":"ping"}`)); err != nil {
//		return err
//	}
//	return nil
//}

func (s *V5WebsocketPrivateService) keepAlive(c *websocket.Conn, cancelFunc context.CancelFunc) {
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
				s.logger.Warn("keep alive timeout")
				cancelFunc()
				return
			}
		}
	}()
}

// Close :
func (s *V5WebsocketPrivateService) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.connection.Close()
}

func (s *V5WebsocketPrivateService) writeMessage(messageType int, body []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.connection.WriteMessage(messageType, body)
}
