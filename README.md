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
  - `docker image prune` (remove unused images)
  - `docker container prune` (remove stopped containers)

## Metadata

This image uses Open Containers Initiative (OCI) labels for metadata.

## Project Structure

- `main.go` — Go web server
- `banners/` — Banner font files
- `static/` — Static assets (CSS, etc.)
- `templates/` — HTML templates
- `Dockerfile` — Docker build instructions
- `.dockerignore` — Files/folders to exclude from image

## Good Practices
- Multi-stage build for minimal image size
- Only standard Go packages used
- Metadata labels included
- Unused Docker objects can be pruned
