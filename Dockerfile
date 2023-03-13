FROM golang:1.20.1-alpine3.17 AS builder
WORKDIR /go/src/consumer-log
COPY . .
RUN \
    apk add protoc protobuf-dev make git && \
    make build

FROM scratch
COPY --from=builder /go/src/consumer-log/consumer-log /bin/consumer-log
ENTRYPOINT ["/bin/consumer-log"]
