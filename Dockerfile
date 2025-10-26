FROM golang:1.24.5-alpine3.21 AS builder
WORKDIR /server
COPY .env ./
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/spirex ./cmd/server

FROM gcr.io/distroless/base-debian12 AS runner
WORKDIR /server
COPY --from=builder /server/bin/spirex .
EXPOSE 8081
CMD ["./spirex"]
