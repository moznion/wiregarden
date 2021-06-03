PROTO_GEN_CONTAINER := "wiregarden-proto-gen:latest"

gen: clean sync
	docker run -it -v $(PWD):/wiregarden -w /wiregarden $(PROTO_GEN_CONTAINER) protoc \
		--plugin=/sbin/protoc-gen-grpc-java \
		--java_out=lib/src/main/java \
		--grpc-java_out=lib/src/main/java \
		protos/*.proto

sync:
	cp ../../../protos/* ./protos/

clean:
	rm -rf ./lib/src/main/java/net/moznion/wiregarden/*.java
	rm ./protos/*.proto

container4protogen:
	docker build . -f Dockerfile -t $(PROTO_GEN_CONTAINER)

