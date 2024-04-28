FROM docker.io/golang:1.22.2-bookworm as builder
WORKDIR /app
COPY . .
RUN : \
    && apt-get update \
    && DEBIAN_FRONTEND=noninteractive \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && go mod tidy \
    && go build -v -o app

FROM docker.io/debian:bookworm-slim
RUN : \
    && apt-get update \
    && DEBIAN_FRONTEND=noninteractive \
    && apt-get install -y ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/app /app/app
CMD ["/app/app"]
