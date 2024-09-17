# syntax=docker/dockerfile:1

FROM golang:latest as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

FROM alpine 

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /server
# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

ENTRYPOINT ["/server"]
