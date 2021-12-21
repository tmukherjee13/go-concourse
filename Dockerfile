FROM golang:alpine as builder
WORKDIR /go/src
COPY . /go/src/
RUN go build -o /go/src/build/check ./cmd/check && \
    go build -o /go/src/build/in ./cmd/in && \
    go build -o /go/src/build/out ./cmd/out

FROM alpine:3.15
COPY --from=builder /go/src/build/ /opt/resource/


