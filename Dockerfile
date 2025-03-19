# Use the official Golang image as the base image
FROM golang:1.23.4-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o telegram-bot-with-go main.go

# Expose the port the application runs on
EXPOSE 8080

# Set environment variables for the application
ENV CONFIG_PATH=/app/config.yaml

# Run the Go application
CMD ["./telegram-bot-with-go"]