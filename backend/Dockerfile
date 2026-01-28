FROM golang:1.21-alpine
WORKDIR /app

# Copy dependency files first for better caching
COPY backend/go.mod ./
# If you have a go.sum file, uncomment the next line
# COPY backend/go.sum ./

RUN go mod download

# Copy the rest of the source code
COPY backend/ .

# Build the application
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]