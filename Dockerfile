# syntax=docker/dockerfile:1

# --- Build stage ---
FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.title="ASCII Art Web"

WORKDIR /app

# Copy go source
COPY . .

# Build the Go binary statically
RUN CGO_ENABLED=0 go build -o ascii-art-web main.go

# --- Run stage ---
FROM scratch

LABEL org.opencontainers.image.title="ASCII Art Web"
LABEL org.opencontainers.image.description="A simple Go web server for ASCII art"
LABEL org.opencontainers.image.authors="Your Name <your.email@example.com>"

WORKDIR /app

COPY --from=builder /app/ascii-art-web .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/banners ./banners

EXPOSE 8080

ENTRYPOINT ["./ascii-art-web"]
