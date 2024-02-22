LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml


run:
	@go build -o bin/ ./cmd/main.go
	BOOTSTRAP_SERVERS=localhost:9092 PORT=2000 PG_DSN="host=http://localhost port=54322 dbname=pg-token-master user=token-user password=token-password sslmode=disable" ./bin/main