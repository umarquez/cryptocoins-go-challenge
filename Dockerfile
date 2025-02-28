FROM golang:1.23 as builder
WORKDIR /app

# Cache Go modules dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the remaining files and build the application
COPY . .
RUN make build

# Create a minimal final image using Alpine
FROM alpine:latest
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder ./app ./app

# Ensure the binary is executable
RUN chmod +x ./app/bin/app

# list the bin directory
RUN ls -lr ./app/bin/
RUN cd /app/bin/
RUN pwd

# Run the application
CMD ["app/bin/app"]
