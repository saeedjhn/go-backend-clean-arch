FROM golang:latest

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# Copy local code to the container image.
WORKDIR /app

RUN mkdir /build

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Install air
#RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/air-verse/air@latest

# Expose port 8000 to the outside world
EXPOSE 8000

# Run
#ENTRYPOINT ["make"]
CMD ["air", "-c", ".air.toml"]
