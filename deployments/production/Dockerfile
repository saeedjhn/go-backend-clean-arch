FROM golang:1.23.2 AS builder

ENV PROJECT_DIR=/app \
    BUILD_DIR=/build \
    GO111MODULE=on \
    CGO_ENABLED=0

# Copy local code to the container image.
WORKDIR $PROJECT_DIR

RUN mkdir $BUILD_DIR

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Compile the application binary
RUN go build -o $BUILD_DIR/main /app/cmd/main.go

# Install air
RUN go install github.com/air-verse/air@latest

# For Production
FROM golang:1.23.2-alpine3.20 AS prod

ENV ROOT_DIR=/ \
    PROJECT_DIR=/app \
    BUILD_DIR=/build \
    GO111MODULE=on \
    CGO_ENABLED=0

# Copy local code to the container image. (current directory)
WORKDIR $PROJECT_DIR

## Copy the Go modules cache from the builder stage to the development stage
COPY --from=builder /go/pkg/mod /go/pkg/mod

# Copy the build directory from the builder stage to the production stage
COPY --from=builder $BUILD_DIR $BUILD_DIR
COPY --from=builder $PROJECT_DIR/configs $PROJECT_DIR/configs

# Expose ports to the outside world
EXPOSE 8000 50000

# Run
# Specify the command to run when the container starts (execute the built binary)
ENTRYPOINT ["/build/main"]