# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git for fetch (if needed) and certs
RUN apk add --no-cache git ca-certificates

# Copy go mod and sum files
COPY . .

# Run mod tidy after copying source
RUN go mod tidy

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o pwgen cmd/server/main.go

# Final Stage
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/pwgen /pwgen

ENV PORT=5069
ENV GIN_MODE=release

EXPOSE 5069

ENTRYPOINT ["/pwgen"]
