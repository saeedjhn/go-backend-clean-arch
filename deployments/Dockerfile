FROM golang:alpine3.19

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# Setup folders
#RUN mkdir /app
WORKDIR /app

RUN mkdir /build

#RUN go mod download
#RUN go install

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

#RUN go install

RUN go get github.com/githubnemo/CompileDaemon

RUN go install github.com/githubnemo/CompileDaemon

EXPOSE 5000

ENTRYPOINT CompileDaemon -directory="." -build="go build -o /build/app" -command="/build/app"
# CMD ["make", "rerun-dev"]
