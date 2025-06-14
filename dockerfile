# Build stage
FROM golang:1.23.5 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o app

# Final image
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8123

CMD ["/app/app"]
