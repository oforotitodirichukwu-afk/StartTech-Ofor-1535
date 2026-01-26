# Use official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# 1. Copy the dependency files
COPY go.mod ./
# If you created a go.sum, copy it too; if not, the next step creates it
COPY go.sum* ./ 

# 2. DOWNLOAD dependencies inside the Docker image
# This fixes the "missing go.sum" error in your images
RUN go mod download

# 3. Copy the rest of the source code
COPY . .

# 4. Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]