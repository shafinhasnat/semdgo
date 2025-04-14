# SEMDGO
![](semdgo.png)
SEMDGO (sem-dee-go) is a high-performance markdown file server written in Go. The name stands for "SErve MarkDown with GO". This project provides a lightweight, efficient solution for serving markdown content with minimal configuration.

## Overview

SEMDGO is designed for developers and content creators who need a simple yet powerful way to serve markdown-based documentation, blogs, or any markdown content. It eliminates the need for complex frontend frameworks or HTML templating while maintaining clean, readable content structure.

## Technical Specifications

- **Content Directory**: `/var/semdgo/content/`
- **Default Entry Point**: `README.md`
- **Port**: 80
- **Architecture Support**: Multi-architecture (amd64, arm64, arm/v7)
- **Runtime**: Containerized (Docker)

## Quick Start

### Docker Deployment

1. **Custom Image Build**:
```Dockerfile
FROM shafinhasnat/semdgo
COPY ./content/ /var/semdgo/content/
CMD ["./semdgo"]
```

2. **Docker Compose Deployment**:
```yaml
services:
  semdgo:
    image: shafinhasnat/semdgo
    container_name: semdgo
    ports:
      - "80:80"
    volumes:
      - ./README.md:/var/semdgo/content/README.md
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