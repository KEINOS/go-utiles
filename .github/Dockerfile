# =============================================================================
#  Test Container for Vaious Go Versions
# =============================================================================
# Default version
ARG VARIANT="1.14-alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

WORKDIR /workspaces

COPY ./go.mod /workspaces/go.mod
COPY ./go.sum /workspaces/go.sum

RUN apk add --no-cache \
    git \
    alpine-sdk \
    build-base \
    && \
    # pre-download go modules (issue #12)
    go mod download -x

ENTRYPOINT go mod download && go test ./...
