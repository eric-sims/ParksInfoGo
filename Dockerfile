FROM golang:1.23-alpine AS builder

# Set the GOARCH to amd64 (x86 architecture)
ARG GOARCH=amd64

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the binary
COPY . .
RUN GOARCH=$GOARCH go build -o server .

# Stage 2: Run the Go server in a minimal container
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080

CMD ["./server"]