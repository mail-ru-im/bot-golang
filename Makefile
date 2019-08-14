all: generate test

$(GOPATH)/bin/easyjson:
	go build -o $(GOPATH)/bin/easyjson github.com/mailru/easyjson/easyjson

.PHONY: test
test:
	go test -v --cover ./...

.PHONY: generate
generate: $(GOPATH)/bin/easyjson
	go generate