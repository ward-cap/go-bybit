GO ?= go
GOCACHE ?= /tmp/go-build
EASYJSON ?= $(GO) run github.com/mailru/easyjson/easyjson

EASYJSON_FILES := $(shell rg -l '`json:' -g '*.go' -g '!*_test.go' -g '!*_easyjson.go')

.PHONY: easyjson regenerate-easyjson clean-easyjson test

easyjson:
	GOCACHE=$(GOCACHE) $(EASYJSON) -all $(EASYJSON_FILES)
	gofmt -w $(EASYJSON_FILES) *_easyjson.go

regenerate-easyjson: clean-easyjson easyjson

clean-easyjson:
	rm -f *_easyjson.go

test:
	GOCACHE=$(GOCACHE) $(GO) test ./...
