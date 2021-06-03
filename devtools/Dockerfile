FROM adoptopenjdk:11.0.11_9-jdk-hotspot-focal

ARG go_tar_file_name="go1.16.4.linux-amd64.tar.gz"

ENV TZ=UTC
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash -
# hadolint ignore=DL3008
RUN DEBIAN_FRONTEND="noninteractive" apt-get -y update \
  && DEBIAN_FRONTEND="noninteractive" apt-get install -y --no-install-recommends \
    nodejs \
    make \
    protobuf-compiler \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*
RUN curl -L -o ${go_tar_file_name} https://golang.org/dl/${go_tar_file_name} \
  && rm -rf /usr/local/go \
  && tar -C /usr/local -xzf ${go_tar_file_name} \
  && rm -f ${go_tar_file_name}
ENV PATH $PATH:/usr/local/go/bin:/root/go/bin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0 \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
RUN npm install -g npm && npm install -g grpc-tools@1.11.1
RUN curl -L -o /sbin/protoc-gen-grpc-java https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/1.38.0/protoc-gen-grpc-java-1.38.0-linux-x86_64.exe \
  && chmod 755 /sbin/protoc-gen-grpc-java

