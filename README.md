# ASCII Art Web Dockerized

A containerized ASCII art web application built with Go, featuring optimized Docker builds and security best practices.

## Quick Start

### Build the Docker image

```bash
docker build -t ascii-art-web:latest .
```

### Run the container

```bash
docker run --rm -p 8080:8080 ascii-art-web:latest
```



## Docker Features

- **Multi-stage build** - Optimized for smaller final image size
- **Distroless base image** - Enhanced security with minimal attack surface
- **Non-root user** - Runs as non-privileged user for better security
- **Layer caching** - Optimized build process with Go modules caching
- **Static binary** - Self-contained executable with no external dependencies


