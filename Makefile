.EXPORT_ALL_VARIABLES:
BIN_SERVER := "./bin/server"
BIN_CLIENT_LIN := "./bin/cl-lin"
BIN_CLIENT_MAC := "./bin/cl-mac"
BIN_CLIENT_WIN := "./bin/cl-win.exe"
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X github.com/brisk84/gophkeeper/pkg/ver.release="develop" -X github.com/brisk84/gophkeeper/pkg/ver.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X github.com/brisk84/gophkeeper/pkg/ver.gitHash=$(GIT_HASH)

build:
	rm -f ./bin/cl-lin ./bin/cl-mac ./bin/cl-win.exe
	GOOS=linux GOARCH=amd64 go build -v -o $(BIN_CLIENT_LIN) -ldflags="$(LDFLAGS)" ./cmd/client/main.go
	GOOS=darwin GOARCH=amd64 go build -v -o $(BIN_CLIENT_MAC) -ldflags="$(LDFLAGS)" ./cmd/client/main.go
	GOOS=windows GOARCH=amd64 go build -v -o $(BIN_CLIENT_WIN) -ldflags="$(LDFLAGS)" ./cmd/client/main.go

generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/gophkeeper.proto

cert:
	cd cert; ./gen.sh; cd ..

run:
	go run ./cmd/server/main.go -a :4343

test:
	go test -v -covermode=count ./...

.PHONY: cert