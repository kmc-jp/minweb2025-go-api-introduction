# Use the official Golang image to build the application
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o hello .

# Use a minimal image to run the application
FROM gcr.io/distroless/base-debian10

# Set the working directory
WORKDIR /

# Copy the built binary
COPY --from=builder /app/hello .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./hello"]
