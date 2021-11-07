all: generate test

GOPATH := $(shell go env GOPATH)

$(GOPATH)/bin/easyjson:
	go build -mod mod -o $(GOPATH)/bin/easyjson github.com/mailru/easyjson/easyjson

$(GOPATH)/bin/golangci-lint:
	go build -mod mod -o $(GOPATH)/bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: test
test:
	go test -v --cover -coverprofile=cover.out ./...

.PHONY: lint
lint: $(GOPATH)/bin/golangci-lint
	$(GOPATH)/bin/golangci-lint run

.PHONY: generate
generate: $(GOPATH)/bin/easyjson
	go generate
