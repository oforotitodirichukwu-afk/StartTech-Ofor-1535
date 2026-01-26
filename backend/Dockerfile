# Use official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy the module files first
COPY go.mod ./ 
COPY go.sum* ./ 

# Download dependencies inside Docker
RUN go mod tidy
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["./main"]
