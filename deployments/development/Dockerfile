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

# Install air
RUN go install github.com/air-verse/air@latest

# For Development
FROM golang:1.23.2-alpine3.20 AS dev

ENV PROJECT_DIR=/app \
    BUILD_DIR=/build \
    GO111MODULE=on \
    CGO_ENABLED=0

# Copy local code to the container image. (current directory)
WORKDIR $PROJECT_DIR

# --from=builder: This instructs Docker to copy files from the builder stage (defined earlier in the Dockerfile using AS builder).
# /app: This refers to the /app directory in the builder stage, which contains your source code and possibly other files.
# .: This copies the files into the current working directory (WORKDIR), which in this case is set to /app in the dev stage.
# COPY --from=builder target source
COPY --from=builder $PROJECT_DIR .

# Copy air binary from Go bin in builder stage
COPY --from=builder /go/bin/air /usr/local/bin/air

# Expose ports to the outside world
EXPOSE 8000 50000

# Run
#ENTRYPOINT ["make"]
CMD ["air", "-c", ".air.toml"]

# For Production