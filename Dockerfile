# Build stage
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/sabify ./cmd/web

# Run stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates

RUN addgroup -S sabify && adduser -S sabify -G sabify

WORKDIR /app

COPY --from=builder /app/sabify /app/sabify

COPY --from=builder /app/ui /app/ui
COPY --from=builder /app/migrations /app/migrations

RUN chown -R sabify:sabify /app

USER sabify

EXPOSE 4000

ENTRYPOINT ["/app/sabify"]
