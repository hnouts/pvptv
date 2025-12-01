FROM golang:1.23-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the server binary
RUN go build -o /server ./cmd/server

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary and required runtime assets
COPY --from=builder /server ./server
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/docs/web ./docs/web

EXPOSE 8080

CMD ["./server"]
