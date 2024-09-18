# ebrick command line interface

The ebrick CLI is a powerful command-line tool that simplifies the development, testing, and deployment of applications built with the eBrick framework. Designed with flexibility in mind, this CLI supports creating modular applications that can transform from monolithic structures to microservices and vise-versa.

## Installation

Before you can use the ebrick CLI, you need to install it. Follow these steps:

### Using Go

```bash
go install github.com/trinitytechnology/ebrick-cli/cmd/ebrick@latest
```

### Download the ebrick CLI binary

- For Linux:

```bash
curl -s https://api.github.com/repos/trinitytechnology/ebrick-cli/releases/latest \
| grep "browser_download_url.*ebrick-linux-amd64" \
| cut -d '"' -f 4 \
| xargs -n 1 curl -L -o ebrick

chmod +x ebrick
mv ./ebrick /usr/local/bin/ebrick
ebrick version
```

- For Windows/MacOS:
Download the binary from https://github.com/trinitytechnology/ebrick-cli/releases.

### Using Docker image

```bash
docker run -it trinitytechnology/ebrick new app
```

### Create new application

```bash
ebrick new app
```

### Create new module

```bash
ebrick new module
```
