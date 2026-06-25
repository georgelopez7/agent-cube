FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/scripts/seed-db
RUN CGO_ENABLED=0 GOOS=linux go build -o /seed-db .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl
WORKDIR /root/
COPY --from=builder /seed-db .
CMD ["./seed-db"]
