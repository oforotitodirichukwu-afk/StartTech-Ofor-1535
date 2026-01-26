# Use official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# 1. Copy only the module file first
COPY go.mod ./

# 2. This is the fix: Force tidy and download INSIDE the container
# This generates the missing go.sum entries the error is complaining about
RUN go mod tidy
RUN go mod download

# 3. Now copy the rest of your source code
COPY . .

# 4. Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]