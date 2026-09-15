package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestCallAPICreatesChildSpanAndLogsContext(t *testing.T) {
	spanRecorder := tracetest.NewSpanRecorder()
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(spanRecorder))
	prevProvider := otel.GetTracerProvider()
	otel.SetTracerProvider(tp)
	defer otel.SetTracerProvider(prevProvider)

	observedCore, recordedLogs := observer.New(zap.DebugLevel)
	logger := zap.New(observedCore).Sugar()

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			body, _ := json.Marshal(map[string]any{
				"retCode": 0,
				"retMsg":  "OK",
				"result":  map[string]any{},
				"time":    1,
			})
			return &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body:    io.NopCloser(strings.NewReader(string(body))),
				Request: req,
			}, nil
		}),
	}

	client := NewClient().
		WithBaseURL("https://api.bybit.test").
		WithHTTPClient(httpClient).
		WithLogger(logger).
		WithAuth("test-api-key", "super-secret").
		withCheckResponseBody(checkV5ResponseBody)

	parentCtx, parentSpan := otel.Tracer("test").Start(context.Background(), "parent")
	parentTraceID := parentSpan.SpanContext().TraceID().String()
	parentSpanID := parentSpan.SpanContext().SpanID().String()

	var dst struct {
		CommonV5Response
		Result map[string]any `json:"result"`
	}

	err := client.postV5JSON(parentCtx, "/v5/order/create", []byte(`{"symbol":"BTCUSDT"}`), "V5OrderService", &dst)
	parentSpan.End()
	if err != nil {
		t.Fatalf("postV5JSON returned error: %v", err)
	}

	var childSpan sdktrace.ReadOnlySpan
	for _, span := range spanRecorder.Ended() {
		if span.Name() == "bybit.V5OrderService" {
			childSpan = span
			break
		}
	}
	if childSpan == nil {
		t.Fatalf("expected child span to be recorded")
	}
	if got := childSpan.SpanContext().TraceID().String(); got != parentTraceID {
		t.Fatalf("trace id mismatch: got %s want %s", got, parentTraceID)
	}
	if got := childSpan.SpanContext().SpanID().String(); got == parentSpanID {
		t.Fatalf("child span reused parent span id: %s", got)
	}

	entries := recordedLogs.AllUntimed()
	if len(entries) < 2 {
		t.Fatalf("expected request and response logs, got %d", len(entries))
	}

	requestFields := entries[0].ContextMap()
	if requestFields["ctx"] == nil {
		t.Fatalf("expected ctx field in request log")
	}
	if got := requestFields["bybit.service"]; got != "V5OrderService" {
		t.Fatalf("unexpected bybit.service: %v", got)
	}
	if got := requestFields["trace_id"]; got != parentTraceID {
		t.Fatalf("unexpected trace_id in request log: %v", got)
	}
	if got := requestFields["span_id"]; got == parentSpanID {
		t.Fatalf("request log reused parent span_id: %v", got)
	}
}

func TestCallAPICreatesRootSpanWithoutParent(t *testing.T) {
	spanRecorder := tracetest.NewSpanRecorder()
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(spanRecorder))
	prevProvider := otel.GetTracerProvider()
	otel.SetTracerProvider(tp)
	defer otel.SetTracerProvider(prevProvider)

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body:    io.NopCloser(strings.NewReader(`{"ret_code":0,"ret_msg":"OK","ext_code":"","ext_info":"","time_now":"1"}`)),
				Request: req,
			}, nil
		}),
	}

	client := NewClient().
		WithBaseURL("https://api.bybit.test").
		WithHTTPClient(httpClient)

	var dst map[string]any
	if err := client.getPubliclyCtx(context.Background(), "/v3/public/time", nil, "TimeService", &dst); err != nil {
		t.Fatalf("getPubliclyCtx returned error: %v", err)
	}

	var rootSpan sdktrace.ReadOnlySpan
	for _, span := range spanRecorder.Ended() {
		if span.Name() == "bybit.TimeService" {
			rootSpan = span
			break
		}
	}
	if rootSpan == nil {
		t.Fatalf("expected root span to be recorded")
	}
	if !rootSpan.SpanContext().IsValid() {
		t.Fatalf("expected valid span context")
	}
	if rootSpan.Parent().IsValid() {
		t.Fatalf("expected root span without parent")
	}
}

func TestCallAPISanitizesSensitiveDataInLogs(t *testing.T) {
	observedCore, recordedLogs := observer.New(zap.DebugLevel)
	logger := zap.New(observedCore).Sugar()

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusUnauthorized,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body:    io.NopCloser(strings.NewReader(`{"retCode":10001,"retMsg":"bad auth"}`)),
				Request: req,
			}, nil
		}),
	}

	client := NewClient().
		WithBaseURL("https://api.bybit.test").
		WithHTTPClient(httpClient).
		WithLogger(logger).
		WithAuth("test-api-key", "super-secret").
		withCheckResponseBody(checkV5ResponseBody)

	var dst map[string]any
	err := client.postV5JSON(
		context.Background(),
		"/v5/order/create",
		[]byte(`{"api_key":"body-api-key","secret":"body-secret","sign":"body-signature","symbol":"BTCUSDT"}`),
		"V5OrderService",
		&dst,
	)
	if err == nil {
		t.Fatalf("expected authentication error")
	}

	entries := recordedLogs.AllUntimed()
	if len(entries) < 3 {
		t.Fatalf("expected request, response and error logs, got %d", len(entries))
	}

	requestFields := entries[0].ContextMap()
	requestHeaders := fmt.Sprintf("%v", requestFields["http.request.header"])
	if strings.Contains(requestHeaders, "test-api-key") {
		t.Fatalf("api key was not masked: %s", requestHeaders)
	}
	if !strings.Contains(requestHeaders, "[REDACTED]") {
		t.Fatalf("signature was not redacted: %s", requestHeaders)
	}

	requestBody, ok := requestFields["http.request.body"].(string)
	if !ok {
		t.Fatalf("unexpected request body type: %T", requestFields["http.request.body"])
	}
	for _, sensitive := range []string{"body-secret", "body-signature", "body-api-key", "super-secret"} {
		if strings.Contains(requestBody, sensitive) {
			t.Fatalf("request body contains sensitive value %q: %s", sensitive, requestBody)
		}
	}
	if !strings.Contains(requestBody, "[REDACTED]") || !strings.Contains(requestBody, "********-key") {
		t.Fatalf("request body was not sanitized: %s", requestBody)
	}

	errorFields := entries[len(entries)-1].ContextMap()
	if errorFields["ctx"] == nil {
		t.Fatalf("expected ctx field in error log")
	}
	if _, ok := errorFields["error"]; !ok {
		t.Fatalf("expected error field in error log")
	}
}
