# Use Go 1.22
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./

# Force Go to use the specified version
ENV GOTOOLCHAIN=local

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Copy binary and config
COPY --from=builder /app/main .
COPY --from=builder /app/.env* ./

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]