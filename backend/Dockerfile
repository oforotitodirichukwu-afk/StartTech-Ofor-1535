# Use official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# 1. Copy go.mod and go.sum first to leverage Docker caching
COPY go.mod ./
# If you don't have a go.sum yet, Docker might complain, 
# so we run tidy right after copying the mod file.

# 2. Download dependencies inside the container
RUN go mod tidy
RUN go mod download

# 3. Copy the rest of the source code
COPY . .

# 4. Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]