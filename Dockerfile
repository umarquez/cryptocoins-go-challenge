FROM golang:1.23 as builder
WORKDIR /app

# Cache Go modules dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the remaining files and build the application
COPY . .
RUN make build

# Run the application
CMD ["./bin/app"]
