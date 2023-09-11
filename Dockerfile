FROM docker.io/golang:1.21.1-bookworm as builder
WORKDIR /app
COPY . .
RUN go build -v -o app

FROM docker.io/debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/app /app/app
CMD ["/app/app"]
