# Use the official Golang image to create a build artifact.
FROM golang:1.19 as builder

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set up project structure in Docker
WORKDIR /app

# Copy go.mod and sum files
COPY ./src/go.mod ./src/go.sum ./src/
COPY ./build/generated/go.mod ./build/generated/

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN cd src && go mod download
RUN cd build/generated && go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src /app/src
COPY ./build /app/build

# Build the Go app
RUN cd src && go build -o ../build/bityagi .

# Start from the latest golang base image
FROM golang:1.17

WORKDIR /app

# Copy the entire spec folder to /app/spec
COPY ./spec /app/spec

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/build/bityagi .
COPY --from=builder /app/src/* /app/src/

# Command to run the executable
ENTRYPOINT ["./bityagi"]
