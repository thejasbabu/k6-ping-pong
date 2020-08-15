FROM golang:1.13-alpine as builder
RUN mkdir -p /k6-ping-pong
COPY . /k6-ping-pong
WORKDIR /k6-ping-pong
RUN go build

FROM alpine
COPY --from=builder /k6-ping-pong/k6-ping-pong ./k6-ping-pong
CMD ./k6-ping-pong