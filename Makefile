.PHONY: proto

WG_PROTO_GEN_CONTAINER := "wiregarden-proto-gen:latest"

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

