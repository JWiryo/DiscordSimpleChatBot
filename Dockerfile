
FROM golang:alpine AS builder

RUN apk --no-cache add ca-certificates

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY ./ .

# Build the application
RUN go build -o DiscordBot .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/DiscordBot .

# Copy binary from this workspace constant folder to docker image's ./constants folder
COPY /constants ./constants

# Export necessary port
EXPOSE 3000

# Build a small image using scratch image from docker
FROM scratch


# Copy SSL Certificates, Go App Binary (DiscordBot) and constants folder to scratch image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /dist/DiscordBot /
COPY --from=builder /dist/constants /constants


# Command to run
ENTRYPOINT ["/DiscordBot"]