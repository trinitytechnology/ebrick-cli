# Use the official Golang image from the Docker Hub
FROM golang:1.22.5-alpine3.20

# Install Git
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Define build argument for version
ARG VERSION=dev

# Build the Go app with the version flag
RUN go build -ldflags "-X main.version=${VERSION}" -o /usr/local/bin/ebrick ./cmd/ebrick

# Command to run the executable
ENTRYPOINT ["ebrick"]
