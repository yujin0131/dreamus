FROM golang:1.24 AS builder

WORKDIR /workspace

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd/main.go ./cmd/main.go
COPY docs ./docs
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o dreamus ./cmd/main.go

FROM debian:bullseye-slim

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /workspace/dreamus .

EXPOSE 8080
ENTRYPOINT [ "./dreamus" ]
