package bybit

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

const (
	// MainNetBaseURL :
	MainNetBaseURL = "https://api.bybit.com"
	// MainNetBaseURL2 :
	MainNetBaseURL2 = "https://api.bytick.com"

	bybitPackageName = "bybit"
	otelTracerName   = "github.com/ward-cap/go-bybit"
)

// Client :
type Client struct {
	httpClient *http.Client

	debug  bool
	logger *zap.SugaredLogger

	baseURL string
	key     string
	secret  string

	referer string

	checkResponseBody        checkResponseBodyFunc
	syncTimeDeltaNanoSeconds int64
}

type apiCall struct {
	service string
	req     *http.Request
	body    []byte
	dst     any
}

// NewClient :
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},

		baseURL:           MainNetBaseURL,
		checkResponseBody: checkResponseBody,
	}
}

// WithHTTPClient :
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient

	return c
}

// WithDebug :
func (c *Client) WithDebug(debug bool) *Client {
	c.debug = debug

	return c
}

// WithLogger :
func (c *Client) WithLogger(logger *zap.SugaredLogger) *Client {
	c.debug = true
	c.logger = logger

	return c
}

// WithAuth :
func (c *Client) WithAuth(key string, secret string) *Client {
	c.key = key
	c.secret = secret

	return c
}

func (c Client) withCheckResponseBody(f checkResponseBodyFunc) *Client {
	c.checkResponseBody = f

	return &c
}

// WithBaseURL :
func (c *Client) WithBaseURL(url string) *Client {
	c.baseURL = url

	return c
}

func (c *Client) WithReferer(referer string) *Client {
	c.referer = referer

	return c
}

// Sign returns an HMAC-SHA256 signature for the provided payload using the client's secret.
func (c *Client) Sign(payload string) string {
	return signPayload(c.secret, payload)
}

// Request :
func (c *Client) Request(req *http.Request, dst any) (err error) {
	return c.callAPI(req.Context(), apiCall{
		service: "Client.Request",
		req:     req,
		body:    cloneRequestBody(req),
		dst:     dst,
	})
}

// hasAuth : check has auth key and secret
func (c *Client) hasAuth() bool {
	return c.key != "" && c.secret != ""
}

func (c *Client) populateSignature(src url.Values) url.Values {
	if src == nil {
		src = url.Values{}
	}

	src.Add("api_key", c.key)
	src.Add("timestamp", strconv.FormatInt(c.getTimestamp(), 10))
	if c.referer != "" {
		src.Add("referer", c.referer)
	}
	src.Add("sign", getSignature(src, c))

	return src
}

func (c *Client) populateSignatureForBody(src []byte) []byte {
	body := map[string]interface{}{}
	if err := json.Unmarshal(src, &body); err != nil {
		panic(err)
	}

	body["api_key"] = c.key
	body["timestamp"] = strconv.FormatInt(c.getTimestamp(), 10)
	if c.referer != "" {
		body["referer"] = c.referer
	}
	body["sign"] = getSignatureForBody(body, c)

	result, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return result
}

func getV5Signature(
	timestamp int64,
	key string,
	queryString string,
	signer interface{ Sign(string) string },
) string {
	return signer.Sign(fmt.Sprintf("%d%s%s", timestamp, key, queryString))
}

func getV5SignatureForBody(
	timestamp int64,
	key string,
	body []byte,
	signer interface{ Sign(string) string },
) string {
	val := strconv.FormatInt(timestamp, 10) + key
	val = val + string(body)
	return signer.Sign(val)
}

func getSignature(src url.Values, signer interface{ Sign(string) string }) string {
	keys := make([]string, len(src))
	i := 0
	_val := ""
	for k := range src {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		_val += k + "=" + src.Get(k) + "&"
	}
	_val = _val[0 : len(_val)-1]
	return signer.Sign(_val)
}

func getSignatureForBody(src map[string]interface{}, signer interface{ Sign(string) string }) string {
	keys := make([]string, len(src))
	i := 0
	_val := ""
	for k := range src {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		_val += k + "=" + fmt.Sprintf("%v", src[k]) + "&"
	}
	_val = _val[0 : len(_val)-1]
	return signer.Sign(_val)
}

func signPayload(secret string, payload string) string {
	h := hmac.New(sha256.New, []byte(secret))
	_, err := io.WriteString(h, payload)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) getPublicly(path string, query url.Values, service string, dst any) error {
	return c.getPubliclyCtx(context.Background(), path, query, service, dst)
}

func (c *Client) getPubliclyCtx(ctx context.Context, path string, query url.Values, service string, dst any) error {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path
	u.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	if err := c.callAPI(ctx, apiCall{
		service: service,
		req:     req,
		dst:     dst,
	}); err != nil {
		return err
	}

	return nil
}

func (c *Client) getV5PrivatelyCtx(ctx context.Context, path string, query url.Values, service string, dst any) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path
	u.RawQuery = query.Encode()

	timestamp := c.getTimestamp()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-BAPI-API-KEY", c.key)
	req.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-BAPI-SIGN", getV5Signature(timestamp, c.key, u.RawQuery, c))

	if err := c.callAPI(ctx, apiCall{
		service: service,
		req:     req,
		dst:     dst,
	}); err != nil {
		return err
	}

	return nil
}

func (c *Client) postV5JSON(ctx context.Context, path string, body []byte, service string, dst any) error {
	if !c.hasAuth() {
		return fmt.Errorf("this is private endpoint, please set api key and secret")
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}
	u.Path = path

	timestamp := c.getTimestamp()
	sign := getV5SignatureForBody(timestamp, c.key, body, c)

	if ctx == nil {
		ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-BAPI-API-KEY", c.key)
	req.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-BAPI-SIGN", sign)
	if c.referer != "" {
		req.Header.Set("X-Referer", c.referer)
	}

	if err := c.callAPI(ctx, apiCall{
		service: service,
		req:     req,
		body:    body,
		dst:     dst,
	}); err != nil {
		return err
	}

	return nil
}

func (c *Client) callAPI(ctx context.Context, call apiCall) (err error) {
	if call.req == nil {
		return errors.New("request should not be nil")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if call.service == "" {
		call.service = "Client"
	}

	serverAddress := ""
	urlFull := ""
	if call.req.URL != nil {
		serverAddress = call.req.URL.Hostname()
		urlFull = sanitizeURL(call.req.URL)
	}

	ctx, span := otel.Tracer(otelTracerName).Start(
		ctx,
		fmt.Sprintf("bybit.%s", call.service),
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("bybit.package", bybitPackageName),
			attribute.String("bybit.service", call.service),
			attribute.String("http.method", call.req.Method),
			attribute.String("url.full", urlFull),
			attribute.String("server.address", serverAddress),
		),
	)
	defer span.End()

	req := call.req.Clone(ctx)
	startedAt := time.Now()
	logger := c.requestLogger(ctx, call.service, req)

	c.logRequest(logger, req, call.body)

	resp, err := c.httpClient.Do(req)
	duration := time.Since(startedAt)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		c.logError(logger, nil, nil, duration, err)
		return err
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		c.logError(logger, resp, nil, duration, err)
		return err
	}

	span.SetAttributes(attribute.Int("http.response.status_code", resp.StatusCode))
	c.logResponse(logger, resp, respBody, duration)

	switch {
	case 200 <= resp.StatusCode && resp.StatusCode <= 299:
		if c.checkResponseBody == nil {
			err = errors.New("checkResponseBody func should be set")
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			c.logError(logger, resp, respBody, duration, err)
			return err
		}
		if err := c.checkResponseBody(respBody); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			c.logError(logger, resp, respBody, duration, err)
			return err
		}
		if err := json.Unmarshal(respBody, call.dst); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			c.logError(logger, resp, respBody, duration, err)
			return err
		}
		span.SetStatus(codes.Ok, "")
		return nil
	case resp.StatusCode == http.StatusBadRequest:
		err = fmt.Errorf("%v: Need to send the request with GET / POST (must be capitalized)", ErrBadRequest)
	case resp.StatusCode == http.StatusUnauthorized:
		err = fmt.Errorf("%w: invalid key/secret", ErrInvalidRequest)
	case resp.StatusCode == http.StatusForbidden:
		err = fmt.Errorf("%w: not permitted", ErrForbiddenRequest)
	case resp.StatusCode == http.StatusNotFound:
		err = fmt.Errorf("%w: wrong path", ErrPathNotFound)
	default:
		err = fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
	c.logError(logger, resp, respBody, duration, err)
	return err
}

func (c *Client) shouldStructuredLog() bool {
	return c.logger != nil && c.debug
}

func (c *Client) requestLogger(ctx context.Context, service string, req *http.Request) *zap.SugaredLogger {
	if !c.shouldStructuredLog() || req == nil || req.URL == nil {
		return nil
	}

	fields := []any{
		"bybit.package", bybitPackageName,
		"bybit.service", service,
		"http.method", req.Method,
		"url.full", sanitizeURL(req.URL),
		"server.address", req.URL.Hostname(),
		"ctx", ctx,
	}

	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.IsValid() {
		fields = append(fields,
			"trace_id", spanContext.TraceID().String(),
			"span_id", spanContext.SpanID().String(),
		)
	}

	return c.logger.With(fields...)
}

func (c *Client) logRequest(logger *zap.SugaredLogger, req *http.Request, body []byte) {
	if logger == nil {
		return
	}
	logger.Debugw("bybit HTTP request",
		"http.request.header", sanitizeHeaders(req.Header),
		"http.request.body", sanitizeRequestBody(req, body),
	)
}

func (c *Client) logResponse(logger *zap.SugaredLogger, resp *http.Response, body []byte, duration time.Duration) {
	if logger == nil {
		return
	}
	fields := []any{"event.duration", duration}
	if resp != nil {
		fields = append(fields,
			"http.response.status_code", resp.StatusCode,
			"http.response.header", sanitizeHeaders(resp.Header),
		)
	}
	if len(body) > 0 {
		fields = append(fields, "http.response.body", sanitizeBody(headerValue(resp, "Content-Type"), body))
	}
	logger.Debugw("bybit HTTP response", fields...)
}

func (c *Client) logError(
	logger *zap.SugaredLogger,
	resp *http.Response,
	respBody []byte,
	duration time.Duration,
	err error,
) {
	if logger == nil {
		return
	}
	fields := []any{"event.duration", duration, "error", err}
	if resp != nil {
		fields = append(fields,
			"http.response.status_code", resp.StatusCode,
			"http.response.header", sanitizeHeaders(resp.Header),
		)
	}
	if len(respBody) > 0 {
		fields = append(fields, "http.response.body", sanitizeBody(headerValue(resp, "Content-Type"), respBody))
	}
	logger.Errorw("bybit HTTP error", fields...)
}

func sanitizeRequestBody(req *http.Request, body []byte) string {
	if len(body) == 0 {
		body = cloneRequestBody(req)
	}
	contentType := ""
	if req != nil {
		contentType = req.Header.Get("Content-Type")
	}
	return sanitizeBody(contentType, body)
}

func cloneRequestBody(req *http.Request) []byte {
	if req == nil || req.GetBody == nil {
		return nil
	}
	body, err := req.GetBody()
	if err != nil {
		return nil
	}
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return nil
	}
	return data
}

func headerValue(resp *http.Response, key string) string {
	if resp == nil {
		return ""
	}
	return resp.Header.Get(key)
}

func sanitizeURL(rawURL *url.URL) string {
	if rawURL == nil {
		return ""
	}
	sanitized := *rawURL
	if sanitized.RawQuery == "" {
		return sanitized.String()
	}
	queryValues, err := url.ParseQuery(sanitized.RawQuery)
	if err != nil {
		return sanitized.String()
	}
	sanitized.RawQuery = sanitizeValues(queryValues).Encode()
	return sanitized.String()
}

func sanitizeHeaders(headers http.Header) map[string][]string {
	if headers == nil {
		return nil
	}
	sanitized := make(map[string][]string, len(headers))
	for key, values := range headers {
		out := make([]string, len(values))
		for i, value := range values {
			out[i] = sanitizeKeyValue(key, value)
		}
		sanitized[key] = out
	}
	return sanitized
}

func sanitizeValues(values url.Values) url.Values {
	if values == nil {
		return nil
	}
	sanitized := make(url.Values, len(values))
	for key, vals := range values {
		out := make([]string, len(vals))
		for i, value := range vals {
			out[i] = sanitizeKeyValue(key, value)
		}
		sanitized[key] = out
	}
	return sanitized
}

func sanitizeBody(contentType string, body []byte) string {
	if len(body) == 0 {
		return ""
	}

	switch {
	case strings.Contains(contentType, "application/json"):
		var payload any
		if err := json.Unmarshal(body, &payload); err != nil {
			return string(body)
		}
		sanitized, err := json.Marshal(sanitizeJSONValue("", payload))
		if err != nil {
			return string(body)
		}
		return string(sanitized)
	case strings.Contains(contentType, "application/x-www-form-urlencoded"):
		values, err := url.ParseQuery(string(body))
		if err != nil {
			return string(body)
		}
		return sanitizeValues(values).Encode()
	default:
		return string(body)
	}
}

func sanitizeJSONValue(key string, value any) any {
	switch typed := value.(type) {
	case map[string]any:
		sanitized := make(map[string]any, len(typed))
		for nestedKey, nestedValue := range typed {
			sanitized[nestedKey] = sanitizeJSONValue(nestedKey, nestedValue)
		}
		return sanitized
	case []any:
		sanitized := make([]any, len(typed))
		for i, item := range typed {
			sanitized[i] = sanitizeJSONValue(key, item)
		}
		return sanitized
	case string:
		return sanitizeKeyValue(key, typed)
	default:
		if isSensitiveKey(key) {
			return redactValue(key)
		}
		return value
	}
}

func sanitizeKeyValue(key string, value string) string {
	if !isSensitiveKey(key) {
		return value
	}
	return redactValue(key, value)
}

func redactValue(key string, value ...string) string {
	normalized := normalizeSensitiveKey(key)
	if strings.Contains(normalized, "api") && strings.Contains(normalized, "key") {
		if len(value) == 0 || value[0] == "" {
			return "[MASKED]"
		}
		return maskAPIKey(value[0])
	}
	return "[REDACTED]"
}

func maskAPIKey(value string) string {
	if len(value) <= 4 {
		return "[MASKED]"
	}
	return strings.Repeat("*", len(value)-4) + value[len(value)-4:]
}

func isSensitiveKey(key string) bool {
	normalized := normalizeSensitiveKey(key)
	return strings.Contains(normalized, "apikey") ||
		strings.Contains(normalized, "secret") ||
		strings.Contains(normalized, "signature") ||
		strings.Contains(normalized, "sign") ||
		strings.Contains(normalized, "authorization") ||
		strings.Contains(normalized, "token")
}

func normalizeSensitiveKey(key string) string {
	key = strings.ToLower(strings.TrimSpace(key))
	key = strings.ReplaceAll(key, "-", "")
	key = strings.ReplaceAll(key, "_", "")
	return key
}

func (c *Client) getTimestamp() int64 {
	return (time.Now().UnixNano() - c.syncTimeDeltaNanoSeconds) / 1000000
}
