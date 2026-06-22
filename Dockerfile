FROM golang:1.25.2 AS builder

# создает папку внутри контейнера и все команды далее будут выполнятсья в этой папке
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cli ./cmd/cli

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /app/cli .

EXPOSE 8080

CMD ["./server"]