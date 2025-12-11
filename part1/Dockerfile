# Use the official Golang image to build the application
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and source code
COPY hello.go .

# Build the application
RUN go build -o hello hello.go

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./hello"]
