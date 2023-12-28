ARG BIN_NAME=mastodon-post
ARG BIN_VERSION=<unknown>

FROM golang:1-alpine AS builder
ARG BIN_NAME
ARG BIN_VERSION

RUN update-ca-certificates

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-X main.version=${BIN_VERSION}" -o ./out/${BIN_NAME} .

FROM scratch
ARG BIN_NAME
COPY --from=builder /src/out/${BIN_NAME} /usr/bin/${BIN_NAME}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/usr/bin/mastodon-post"]

LABEL license="MIT"
LABEL org.opencontainers.image.licenses="MIT"
LABEL maintainer="Chris Dzombak <https://www.dzombak.com>"
LABEL org.opencontainers.image.authors="Chris Dzombak <https://www.dzombak.com>"
LABEL org.opencontainers.image.url="https://github.com/cdzombak/mastodon-post"
LABEL org.opencontainers.image.documentation="https://github.com/cdzombak/mastodon-post/blob/main/README.md"
LABEL org.opencontainers.image.source="https://github.com/cdzombak/mastodon-post.git"
LABEL org.opencontainers.image.version="${BIN_VERSION}"
LABEL org.opencontainers.image.title="${BIN_NAME}"
LABEL org.opencontainers.image.description="The simplest possible CLI tool for posting to Mastodon"
