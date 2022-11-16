FROM golang:latest AS builder
WORKDIR /build
COPY . .
RUN go build -o /build/app gitlab.ozon.dev/kunata928/telegramBot/cmd/bot

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /build/app /app
CMD ["/app"]
