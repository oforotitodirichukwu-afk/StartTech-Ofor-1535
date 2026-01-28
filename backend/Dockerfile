FROM golang:1.21-alpine
WORKDIR /app

# Copy the backend source files
COPY . .

# Download dependencies and tidy up
RUN go mod download
RUN go mod tidy

# Build the binary
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]