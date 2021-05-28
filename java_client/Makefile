PROTO_GEN_CONTAINER := "wiregarden-proto-gen:latest"

java: sync
	docker run -it -v $(PWD):/wiregarden -w /wiregarden $(PROTO_GEN_CONTAINER) make _java

_java:
	protoc --java_out=lib/src/main/java protos/*.proto

sync:
	cp ../protos/* ./protos/

container4protogen:
	docker build . -f Dockerfile -t $(PROTO_GEN_CONTAINER)

