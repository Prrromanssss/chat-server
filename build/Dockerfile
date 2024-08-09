FROM golang:1.22.5-alpine AS builder

COPY app /github.com/Prrromanssss/chat-server/source
WORKDIR /github.com/Prrromanssss/chat-server/source

RUN go mod download
RUN go build -o ./bin/chat-server cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/Prrromanssss/chat-server/source/bin/chat-server .

ADD config.yaml /config.yaml

CMD ["./chat-server"]