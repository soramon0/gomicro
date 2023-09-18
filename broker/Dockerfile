FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

ENV CGO_ENABLED=0
ENV PORT=80

RUN go build -o brokerApp ./cmd/api/*

RUN chmod +x /app/brokerApp

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]
