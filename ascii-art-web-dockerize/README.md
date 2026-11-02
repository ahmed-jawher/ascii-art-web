# ASCII Art Web — Dockerize

Stage 3 of [ASCII Art Web](../README.md). It packages the application in a Docker container so it runs the same way on any machine.

## What's New

- Code reorganized into an `internal/` package (`ascii.go`, `handlers.go`) to keep `main.go` small
- A `Dockerfile` based on the official `golang` image that exposes port `8080`
- Runs anywhere Docker is installed, with no local Go setup required

## Run with Docker

```bash
docker build -t ascii-art-web:latest .
docker run --rm -p 8080:8080 ascii-art-web:latest
```

Then open <http://localhost:8080>.

## Run without Docker

```bash
go run main.go
```
