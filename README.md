## SEMDGO
SEMDGO (sem-dee-go) is a tool designed to serve markdown files efficiently. The project was initiated on April 8, 2025, and is actively maintained.

This tool aims to simplify the process of serving markdown-based content, providing an alternative to traditional HTML or frontend frameworks. SEMDGO is ideal for users who prefer writing in markdown and need a straightforward solution for hosting their markdown content.

### Running semdgo
Craete a demo markdown file and name it `README.md`
```yaml
services:
  semdgo:
    image: shafinhasnat/semdgo:0.1
    container_name: semdgo
    ports:
      - "80:80"
    volumes:
      - ./README.md:/content/README.md
```
deploy the docker compose definition with `docker-compose up -d` command.

### Building semdgo from source
To build from source - 
```bash
go build
```
Build the docker image for several cpu architechture with the following command-
```bash
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t shafinhasnat/semdgo --push .
```