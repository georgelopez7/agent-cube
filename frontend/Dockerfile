FROM oven/bun:1-alpine AS builder
WORKDIR /app
COPY package.json bun.lock ./
RUN bun install --frozen-lockfile
COPY . .
ARG VITE_WEBSOCKET_URL=ws://localhost:8080/api/v1/ws
ENV VITE_WEBSOCKET_URL=${VITE_WEBSOCKET_URL}
ARG VITE_SITE_URL=http://localhost:3001
ENV VITE_SITE_URL=${VITE_SITE_URL}
ARG VITE_GITHUB_URL=https://github.com/georgelopez7/agent-cube
ENV VITE_GITHUB_URL=${VITE_GITHUB_URL}
ARG VITE_X_URL=https://x.com/xgeloop
ENV VITE_X_URL=${VITE_X_URL}
RUN bun run build


FROM oven/bun:1-alpine AS runner
WORKDIR /app
COPY --from=builder /app/.output ./.output
EXPOSE 3000
CMD ["bun", "run", ".output/server/index.mjs"]
