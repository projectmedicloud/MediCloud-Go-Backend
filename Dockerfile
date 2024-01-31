# Use the official Go image as a parent image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
# COPY go.mod go.sum ./
# RUN go mod download

# Copy the rest of the application's code
COPY . .

# Build the application
RUN go build -o main ./cmd/service

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
