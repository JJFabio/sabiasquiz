# Build stage
FROM golang:1.26.0-alpine AS builder

WORKDIR /app

# Copy dependencies
COPY go.mod ./
# COPY go.sum ./ # No go.sum yet as we haven't added external deps

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o sabiasquiz ./server

# Run stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/sabiasquiz .
# Copy public assets
COPY --from=builder /app/server/public ./public

# Expose port
EXPOSE 8080

# Run the app
CMD ["./sabiasquiz"]
