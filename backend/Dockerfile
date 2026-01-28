FROM golang:1.21-alpine
WORKDIR /app

# Copy all files from the backend directory
COPY . .

# Ensure dependencies are present
RUN go mod download
RUN go mod tidy

# Build the application binary
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]