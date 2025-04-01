# Use the official Golang image
FROM golang:1.22.3

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source code
COPY . .

# Expose the application port
EXPOSE 3001

# Command to run the Go application
CMD ["go", "run", "./cmd/main.go"]
