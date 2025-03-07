package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"time"
)

const (
	WebsocketBaseURL = "wss://stream.bybit.com"
)

// WebSocketClient :
type WebSocketClient struct {
	debug  bool
	logger *zap.SugaredLogger

	baseURL string
	key     string
	secret  string
}

func (c *WebSocketClient) debugf(format string, v ...interface{}) {
	if c.debug {
		c.logger.Infof(format, v...)
	}
}

// NewWebsocketClient :
func NewWebsocketClient() *WebSocketClient {
	return &WebSocketClient{
		baseURL: WebsocketBaseURL,
	}
}

// WithDebug :
func (c *WebSocketClient) WithDebug(debug bool) *WebSocketClient {
	c.debug = debug

	return c
}

// WithLogger :
func (c *WebSocketClient) WithLogger(logger *zap.SugaredLogger) *WebSocketClient {
	c.debug = true
	c.logger = logger

	return c
}

// WithAuth :
func (c *WebSocketClient) WithAuth(key string, secret string) *WebSocketClient {
	c.key = key
	c.secret = secret

	return c
}

// WithBaseURL :
func (c *WebSocketClient) WithBaseURL(url string) *WebSocketClient {
	c.baseURL = url

	return c
}

// hasAuth : check has auth key and secret
func (c *WebSocketClient) hasAuth() bool {
	return c.key != "" && c.secret != ""
}

func (c *WebSocketClient) buildAuthParam() ([]byte, error) {
	if !c.hasAuth() {
		return nil, fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	expires := time.Now().Unix()*1000 + 10000
	req := fmt.Sprintf("GET/realtime%d", expires)
	s := hmac.New(sha256.New, []byte(c.secret))
	if _, err := s.Write([]byte(req)); err != nil {
		return nil, err
	}
	signature := hex.EncodeToString(s.Sum(nil))
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "auth",
		Args: []interface{}{c.key, expires, signature},
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
