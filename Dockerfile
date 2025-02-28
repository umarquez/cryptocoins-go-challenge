FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN make build
RUN chmod +x ./bin/app
ENTRYPOINT ["./bin/app"]

