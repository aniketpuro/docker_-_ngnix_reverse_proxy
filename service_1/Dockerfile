FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
RUN apk add --no-cache curl
EXPOSE 8001

HEALTHCHECK --interval=10s --timeout=2s --start-period=1s --retries=2 \
  CMD curl --fail http://localhost:8001/ping || exit 1

CMD ["./main"]
