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
COPY --from=builder /build/bin/gloc /bin/
WORKDIR /workdir
ENTRYPOINT ["/bin/gloc"]
