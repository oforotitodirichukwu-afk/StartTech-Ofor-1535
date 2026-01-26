# Use the official Golang image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# 1. Copy the module files first
COPY go.mod ./
# Note: We don't copy go.sum yet because it might be incomplete

# 2. Force a 'tidy' and download inside the container
# This is the "magic fix" for the error in your screenshots
RUN go mod tidy
RUN go mod download

# 3. Copy the rest of the source code
COPY . .

# 4. Build the application
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Run the app
CMD ["./main"]