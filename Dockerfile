# syntax=docker/dockerfile:1

# --- Build stage ---
FROM golang:1.22-alpine AS builder

# Install ca-certificates for HTTPS requests (if needed)
RUN apk --no-cache add ca-certificates git

# Add labels for better image metadata
LABEL org.opencontainers.image.title="ASCII Art Web"
LABEL org.opencontainers.image.description="A simple Go web server for ASCII art stylization"
LABEL org.opencontainers.image.version="1.0.0"


WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum* ./

# Download dependencies (cached if go.mod/go.sum haven't changed)
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the Go binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o ascii-art-web main.go

# --- Runtime stage ---
FROM gcr.io/distroless/static-debian12:nonroot

# Add runtime labels
LABEL org.opencontainers.image.title="ASCII Art Web"
LABEL org.opencontainers.image.description="A simple Go web server for ASCII art stylization"
LABEL org.opencontainers.image.version="1.0.0"


WORKDIR /app

# Copy the binary and required files from builder stage
COPY --from=builder /app/ascii-art-web .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/banners ./banners

# Expose port
EXPOSE 8080

# Use non-root user (distroless nonroot image provides this)
USER nonroot:nonroot

# Set entrypoint
ENTRYPOINT ["./ascii-art-web"]