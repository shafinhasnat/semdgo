# SEMDGO
![](semdgo.png)
SEMDGO (sem-dee-go) is a high-performance markdown file server written in Go. The name stands for "SErve MarkDown with GO". This project provides a lightweight, efficient solution for serving markdown content with minimal configuration.

## Overview

SEMDGO is designed for developers and content creators who need a simple yet powerful way to serve markdown-based documentation, blogs, or any markdown content. It eliminates the need for complex frontend frameworks or HTML templating while maintaining clean, readable content structure.

## Technical Specifications

- **Content Directory**: `/var/semdgo/content/`
- **Default Entry Point**: `README.md`
- **Default Port**: 80
- **Architecture Support**: Multi-architecture (amd64, arm64, arm/v7)
- **Runtime**: Containerized (Docker)

## Configuration

SEMDGO can be configured by placing a `semdgo.json` file in the working directory (next to the binary or mounted into the container).

### Custom Port

```json
{
  "port": 8080
}
```

### Automatic HTTPS with Let's Encrypt

SEMDGO supports automatic TLS certificate provisioning via Let's Encrypt — no manual certificate management required. When enabled, HTTP traffic on port 80 is automatically redirected to HTTPS on port 443.

> **Requirements**: your domain must point to the server and ports 80 and 443 must be publicly reachable.

```json
{
  "letsencrypt": {
    "enabled": true,
    "domain": "docs.example.com",
    "email": "you@example.com"
  }
}
```

Certificates are cached at `/var/semdgo/certs` and auto-renewed before expiry.

If no `semdgo.json` is present, SEMDGO defaults to port 80 over plain HTTP.

## Quick Start

### Docker Deployment

1. **Custom Image Build**:
```Dockerfile
FROM shafinhasnat/semdgo
COPY ./content/ /var/semdgo/content/
CMD ["./semdgo"]
```

2. **Docker Compose Deployment**:

Plain HTTP on a custom port:
```yaml
services:
  semdgo:
    image: shafinhasnat/semdgo
    container_name: semdgo
    ports:
      - "8080:8080"
    volumes:
      - ./README.md:/var/semdgo/content/README.md
      - ./semdgo.json:/semdgo.json
```

With Let's Encrypt HTTPS:
```yaml
services:
  semdgo:
    image: shafinhasnat/semdgo
    container_name: semdgo
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./README.md:/var/semdgo/content/README.md
      - ./semdgo.json:/semdgo.json
      - semdgo_certs:/var/semdgo/certs

volumes:
  semdgo_certs:
```

Deploy using:
```bash
docker-compose up -d
```

## Building from Source

### Local Build
```bash
go build ./cmd/server -o ./dist/semdgo
```

### Multi-architecture Docker Build
```bash
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -t shafinhasnat/semdgo \
  --push .
```