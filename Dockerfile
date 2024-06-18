# build stage
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main ./cmd

# final stage
FROM alpine:3.20.0
WORKDIR /app
COPY --from=builder /app/main /app/

# Create a non-root user and switch to it
RUN adduser -D nonroot && chown -R nonroot:nonroot /app
USER nonroot

CMD ["./main"]
EXPOSE 8080
