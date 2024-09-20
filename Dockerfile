FROM golang:1.23.1-alpine AS builder

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the binary
COPY . .
RUN go build -o server .

# Stage 2: Run the Go server in a minimal container
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080

CMD ["./server"]