# Use official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# 1. Copy the module file
COPY go.mod ./

# 2. This is the FIX for the errors in your screenshots
# It downloads the dependencies and creates the checksums inside Docker
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