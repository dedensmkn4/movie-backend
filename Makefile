
.PHONY: install-tools
install-tools:
	go get -u \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/envoyproxy/protoc-gen-validate

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
	go test -race ./...

.PHONY: test-cover
test-cover:
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

.PHONY: test-lint
lint:
	golint ./...
	go vet ./...

.PHONY: test-integration
test-integration:
	go test -tags integration -v ./users

.PHONY: build
build:
	go build -o bin/grpc_server cmd/grpc-server/*.go
	go build -o bin/http_server cmd/http-server/*.go