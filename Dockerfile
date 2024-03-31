# Dockerfile definition for Backend application service.

# From which image we want to build. This is basically our environment.
# Change golang version from 1.19 ro 1.22
FROM golang:1.22 as Build

# This will copy all the files in our repo to the inside the container at root location.
COPY . .

# Disable CGO to get static linked binary
ENV CGO_ENABLED=0

# Build our binary at root location.
RUN GOPATH= go build -o /main cmd/main.go

####################################################################
# This is the actual image that we will be using in production.
FROM alpine:latest

# We need to copy the binary from the build image to the production image.
COPY --from=Build /main .

# This is the port that our application will be listening on.
EXPOSE 1323

# This is the command that will be executed when the container is started.
ENTRYPOINT ["./main"]