FROM golang:1.24-bullseye AS builder

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
  upx-ucl

WORKDIR /build

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go build \
  -ldflags='-w -s -extldflags "-static"' \
  -o ./bin/gloc cmd/gloc/main.go \
  && upx-ucl --best --ultra-brute ./bin/gloc

FROM scratch

ARG APPLICATION="gloc"
ARG DESCRIPTION="🚀 gloc – A blazing-fast LOC (Lines of Code) counter in Go, inspired by tokei. Simple & efficient!"
ARG PACKAGE="trinhminhtriet/gloc"

LABEL org.opencontainers.image.ref.name="${PACKAGE}" \
  org.opencontainers.image.authors="Triet Trinh <contact@trinhminhtriet.com>" \
  org.opencontainers.image.documentation="https://github.com/${PACKAGE}/README.md" \
  org.opencontainers.image.description="${DESCRIPTION}" \
  org.opencontainers.image.licenses="MIT" \
  org.opencontainers.image.source="https://github.com/${PACKAGE}"

COPY --from=builder /build/bin/gloc /bin/
WORKDIR /workdir
ENTRYPOINT ["/bin/gloc"]
