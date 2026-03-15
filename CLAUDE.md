# CLAUDE.md

## Project Overview

SEMDGO is a lightweight markdown file server written in Go. It serves `.md` files from a configured content directory, rendering them to HTML via a built-in template. Non-markdown files (images, etc.) are served as static assets.

## Repository Structure

```
cmd/server/main.go              — entry point, wires config → handler, starts HTTP/HTTPS
internal/config/config.go       — Config struct, Load() from semdgo.json, Validate()
internal/handler/handler.go     — New(contentPath, entrypoint) returns http.HandlerFunc
internal/renderer/markdown.go   — MDtoHTML() wraps gomarkdown
internal/utils/path.go          — ClickedHyperlink() maps request URL to filesystem path
templates/200                   — Go HTML template for rendered markdown pages
Dockerfile                      — multi-stage build; copies binary + templates into alpine
```

## Build & Run

```bash
# build
go build -o ./dist/semdgo ./cmd/server

# run (reads semdgo.json from cwd, templates/ must be present)
./dist/semdgo

# or run directly
go run ./cmd/server
```

## Configuration (`semdgo.json`)

All fields are optional. Place the file next to the binary (or mount it into the container).

| Key                  | Type   | Default                  | Description                              |
|----------------------|--------|--------------------------|------------------------------------------|
| `port`               | int    | `80`                     | HTTP listen port                         |
| `content_path`       | string | `/var/semdgo/content`    | Directory containing markdown files      |
| `content_entrypoint` | string | `README.md`              | File served at `/`                       |
| `letsencrypt.enabled`| bool   | `false`                  | Enable automatic HTTPS via Let's Encrypt |
| `letsencrypt.domain` | string | —                        | Required when `enabled` is true          |
| `letsencrypt.email`  | string | —                        | Optional, for cert expiry notifications  |

Startup validation (via `config.Validate`) will fatal if:
- `content_path` does not exist or is not a directory
- `content_path/content_entrypoint` does not exist
- `letsencrypt.enabled` is true but `letsencrypt.domain` is empty

## Key Implementation Notes

- **Template location**: `templates/200` is loaded relative to the working directory at request time via `template.ParseFiles`. When running locally the binary must be executed from the repo root (or wherever `templates/` lives). The Dockerfile copies `templates/` to `/templates` at the container root and runs the binary from `/`.
- **Static assets**: any request for a path not ending in `.md` is passed to `http.ServeFile`, so images and other files linked from markdown are served directly from `content_path`.
- **Let's Encrypt**: certs are cached at `/var/semdgo/certs`. HTTP on `:80` is kept alive for the ACME HTTP-01 challenge and redirects all other traffic to HTTPS on `:443`.

## Docker

```bash
# single-arch
docker build -t shafinhasnat/semdgo .

# multi-arch push
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -t shafinhasnat/semdgo --push .
```
