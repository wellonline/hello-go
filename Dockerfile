FROM golang:1.21.5 AS build_base
ENV CGO_ENABLED 0
# Set the Current Working Directory inside the container
WORKDIR /tmp/hello-go

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/hello-go .

# Start fresh from a smaller image
FROM alpine:3.9

COPY --from=build_base /tmp/hello-go/out/hello-go /app/hello-go

# This container exposes port 8080 to the outside world
EXPOSE 80

# Run the binary program produced by `go install`
CMD ["/app/hello-go"]
