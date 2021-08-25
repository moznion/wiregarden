.PHONY: proto test lint sec fmt fix

PKGS := $(shell go list ./...)
WG_PROTO_GEN_CONTAINER := "wiregarden-proto-gen:latest"
GO_BUILD_CONTAINER := "golang:1.17-buster"

check: fmt-check test lint vet sec
check-ci: fmt-check test vet sec

proto:
	docker run -it -v $(PWD):/wiregarden -w /wiregarden $(WG_PROTO_GEN_CONTAINER) make _proto

_proto:
	protoc \
		--go_out=. \
		--go_opt=module=github.com/moznion/wiregarden \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/moznion/wiregarden \
		protos/*.proto

container4protogen:
	docker build . -f devtools/Dockerfile -t $(WG_PROTO_GEN_CONTAINER)

build:
ifndef GOOS
	@echo "[error] \$$GOOS must be specified"
	@exit 1
endif
ifndef GOARCH
	@echo "[error] \$$GOARCH must be specified"
	@exit 1
endif
	docker run -it -v $(PWD):/wiregarden -w /wiregarden \
		-e GOOS=$(GOOS) \
		-e GOARCH=$(GOARCH) \
		$(GO_BUILD_CONTAINER) \
		go mod vendor && \
		go build \
			-ldflags '-X "github.com/moznion/wiregarden/internal.Revision=$(shell git rev-parse HEAD)" -X "github.com/moznion/wiregarden/internal.Version=$(shell git describe --abbrev=0 --tags)"' \
			-o ./bin/wiregarden-server_$(GOOS)_$(GOARCH) ./cmd/wiregarden-server

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

github-docker-login:
ifndef DOCKER_USER
	@echo "[error] \$$DOCKER_USER must be specified"
	@exit 1
endif
ifndef DOCKER_PSWD_FILE
	@echo "[error] \$$DOCKER_PSWD_FILE must be specified"
	@exit 1
endif
	cat $(DOCKER_PSWD_FILE) | docker login ghcr.io --username $(DOCKER_USER) --password-stdin

e2e-docker-container:
	docker build . -f ./devtools/e2etest/Dockerfile -t wiregarden-e2e-test:latest
	docker tag wiregarden-e2e-test:latest ghcr.io/moznion/wiregarden/wiregarden-e2e-test:latest

e2e-docker-push: github-docker-login
	docker push ghcr.io/moznion/wiregarden/wiregarden-e2e-test:latest

