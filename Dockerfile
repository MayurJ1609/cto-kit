# Build Stage
FROM golang:1.21.4-alpine3.18 AS builder

WORKDIR /app
COPY api ./api
COPY cmd ./cmd
COPY internal ./internal
COPY vendor ./vendor
COPY go.mod ./go.mod
COPY go.sum ./go.sum

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/api/main.go

# Final Stage
FROM alpine:3.18.4

# Set a non-root user for added security
RUN addgroup -S skor && adduser -S skor -G skor

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/app .

# Create a directory for serving static files
RUN mkdir -p public

# Install libcap
RUN apk add --no-cache libcap tzdata

# Set the NET_BIND_SERVICE capability and drop other capabilities
RUN setcap 'cap_net_bind_service=+ep' ./app

# Change to the non-root user
USER skor

EXPOSE 80

CMD ["./app"]
