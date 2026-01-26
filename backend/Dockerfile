FROM golang:1.21-alpine

WORKDIR /app

# 1. Copy everything first (including main.go, go.mod)
COPY . .

# 2. NOW run tidy (it will see main.go and find the MongoDB imports)
RUN go mod tidy

# 3. Build the application
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]