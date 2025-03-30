FROM golang:1.23.7-alpine AS builder

WORKDIR /app/
COPY . .

RUN go mod download
RUN go build -o ./bin/chat-auth cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/bin/chat-auth .

CMD ["./chat-auth"]

