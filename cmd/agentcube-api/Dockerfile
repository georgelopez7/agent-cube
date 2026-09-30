# syntax=docker/dockerfile:1

FROM golang:1.27-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /agentcube-api ./cmd/agentcube-api

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build --chown=nonroot:nonroot /agentcube-api ./agentcube-api
EXPOSE 8080
ENTRYPOINT ["/app/agentcube-api"]
