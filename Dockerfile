# Use the official Go image to create a build artifact.
FROM golang:1.21-alpine AS builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Copy go.mod and go.sum to download dependencies
COPY go.* ./
RUN go mod download

# Copy the rest of the application's source code
COPY . .

# Build the binary.
RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/service

# Use a small image to run the application
FROM alpine:3.19
COPY --from=builder /bin/app /bin/app

# Command to run the executable
ENTRYPOINT ["/bin/app"]
