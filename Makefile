.PHONY: proto test lint sec fmt fix

PKGS := $(shell go list ./...)
WG_PROTO_GEN_CONTAINER := "wiregarden-proto-gen:latest"
GO_BUILD_CONTAINER := "golang:1.16.0-alpine3.13"

check: fmt-check test lint vet sec

proto:
	docker run -it -v $(PWD):/wiregarden -w /wiregarden $(WG_PROTO_GEN_CONTAINER) make _proto

_proto:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		grpc/messages/*.proto

container4protogen:
	docker build . -f devtools/Dockerfile -t $(WG_PROTO_GEN_CONTAINER)

build:
	docker run -it -v $(PWD):/wiregarden -w /wiregarden \
		-e GOOS=$(GOOS) \
		-e GOARCH=$(GOARCH) \
		$(GO_BUILD_CONTAINER) \
		go build -o ./bin/wiregarden-server_$(GOOS)_$(GOARCH) ./cmd/wiregarden-server

clean:
	rm -f ./bin/wiregarden-server*

test:
	go test -v $(PKGS)

lint:
	golangci-lint run -v

vet:
	go vet $(PKGS)

sec:
	gosec ./...

fmt-check:
	goimports -l **/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi \

fmt:
	gofmt -w -s **/*.go
	goimports -w **/*.go

lint-fix:
	golangci-lint run -v --fix

fix:
	$(MAKE) fmt
	$(MAKE) lint-fix
