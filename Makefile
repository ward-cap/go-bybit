GO ?= go
GOCACHE ?= /tmp/go-build
EASYJSON ?= $(GO) tool easyjson

# Plain grep, not rg: the CI runner that generates these has no ripgrep.
EASYJSON_FILES := $(shell grep -l '`json:' $(filter-out %_test.go %_easyjson.go,$(wildcard *.go)))

.PHONY: easyjson regenerate-easyjson clean-easyjson test

# Hand-written code calls easyjson.Unmarshal on these types, so the package does
# not compile without generated methods - and easyjson needs it to compile. The
# stubs pass breaks that loop on a clean checkout.
easyjson:
	GOCACHE=$(GOCACHE) $(EASYJSON) -stubs -all $(EASYJSON_FILES)
	GOCACHE=$(GOCACHE) $(EASYJSON) -all $(EASYJSON_FILES)
	gofmt -w $(EASYJSON_FILES) *_easyjson.go

regenerate-easyjson: clean-easyjson easyjson

clean-easyjson:
	rm -f *_easyjson.go

test:
	GOCACHE=$(GOCACHE) $(GO) test ./...
