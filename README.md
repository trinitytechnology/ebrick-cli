# ebrick command line interface

The ebrick CLI is a powerful command-line tool that simplifies the development, testing, and deployment of applications built with the eBrick framework. Designed with flexibility in mind, this CLI supports creating modular applications that can transform from monolithic structures to microservices and vise-versa.

## Installation

Before you can use the ebrick CLI, you need to install it. Follow these steps:

### Using Go

```bash
go install github.com/trinitytechnology/ebrick-cli/cmd/ebrick@latest
```

### Download the ebrick CLI binary

- For macOS/Linux:

```bash
curl -L -o ebrick https://example.com/ebrick/releases/latest/download/ebrick-linux-amd64
```

- For Windows:
Download the binary from https://example.com/ebrick/releases.

### Using Docker image

```bash
docker run -it trinitytechnology/ebrick new app
```
