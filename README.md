# ASCII Art Web Dockerized

## Build the Docker image

```sh
docker build -t ascii-art-web:latest .
```

## Run the container

```sh
docker run --rm -p 8080:8080 ascii-art-web:latest
```

- The `--rm` flag ensures the container is removed after exit (helps with garbage collection).
- To clean up unused images/containers:


## Project Structure

- `main.go` — Go web server
- `banners/` — Banner font files
- `static/` — Static assets (CSS, etc.)
- `templates/` — HTML templates
- `Dockerfile` — Docker build instructions
- `.dockerignore` — Files/folders to exclude from image
