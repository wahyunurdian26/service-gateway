FROM golang:alpine AS builder

WORKDIR /app
COPY . .

WORKDIR /app/gateway
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gateway/main .
CMD ["./main"]
