FROM linuxserver/wireguard:amd64-version-v1.0.20210424

WORKDIR /tmp
RUN curl -LO https://golang.org/dl/go1.16.5.linux-amd64.tar.gz \
  && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin:/root/go/bin
RUN go install golang.org/x/tools/cmd/goimports@v0.1.2 && go install github.com/securego/gosec/v2/cmd/gosec@v2.8.0
RUN apt-get update -y && apt-get install --no-install-recommends -y vim && rm -rf /var/lib/apt/lists/*
WORKDIR /
RUN touch /config/wg0.conf && chmod 600 /config/wg0.conf
LABEL org.opencontainers.image.source = "https://github.com/moznion/wiregarden"

