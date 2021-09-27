.PHONY: generate
generate:
	protoc -I=api/proto \
		-I=third_party \
		--go_out=internal/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal/pb \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=internal/pb \
		--grpc-gateway_opt=paths=source_relative \
		api/proto/*.proto

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: build
build:
	go build -o bin/grpc_server cmd/grpc-server/main.go
	go build -o bin/http_server cmd/http-server/main.go

