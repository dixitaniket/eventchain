# syntax=docker/dockerfile:1

ARG GO_VERSION="1.20"
ARG RUNNER_IMAGE="alpine"

# --------------------------------------------------------
# Builder
# --------------------------------------------------------

FROM golang:${GO_VERSION}-alpine as builder

ARG GIT_VERSION
ARG GIT_COMMIT

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers \
    jq \
    sed

# Download go dependencies
WORKDIR /app
COPY go.mod go.sum ./

# Copy the remaining files
COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go build \
        -mod=readonly \
        -tags "muslc" \
        -trimpath \
        -o /app/build/eventchaind \
        /app/cmd/eventchaind/main.go

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

FROM ${RUNNER_IMAGE}

RUN apk add bash \
    libgcc \
    jq

COPY --from=builder /app/build/eventchaind /usr/bin/eventchaind
COPY --from=builder /app/scripts/init.sh ./init.sh

ENV HOME /
WORKDIR $HOME

# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657
# grpc rpc
EXPOSE 8080

CMD ["./init.sh"]