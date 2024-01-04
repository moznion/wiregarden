FROM linuxserver/wireguard:amd64-version-v1.0.20210914

WORKDIR /tmp
RUN curl -LO https://golang.org/dl/go1.21.5.linux-amd64.tar.gz \
  && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin:/root/go/bin
RUN go install golang.org/x/tools/cmd/goimports@v0.16.1 && go install github.com/securego/gosec/v2/cmd/gosec@v2.18.2
WORKDIR /
RUN apk add g++ && apk add make
RUN mkdir -p /config/wg_confs && touch /config/wg_confs/wg0.conf && chmod 600 /config/wg_confs/wg0.conf
LABEL org.opencontainers.image.source = "https://github.com/moznion/wiregarden"

