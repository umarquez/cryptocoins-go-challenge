FROM golang:1.23 as builder
WORKDIR /app

# Cache Go modules dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the remaining files and build the application
COPY . .
RUN make build

# list the directory tree
RUN ls -lr ./bin

# Create a minimal final image using Alpine
FROM alpine:latest
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/app ./bin/app

# list the directory tree
RUN ls -lr ./bin

# Ensure the binary is executable
RUN chmod +x ./bin/app
ENTRYPOINT ["./bin/app"]