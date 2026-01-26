# Use official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy Go module files first (for caching)
COPY go.mod go.sum ./

# Download dependencies INSIDE the container
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./main"]
