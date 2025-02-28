FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build
RUN chmod +x ./bin/app
ENTRYPOINT ["./bin/app"]

