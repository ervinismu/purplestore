FROM golang:1.20.0-alpine3.17 AS builder

# Set working directory
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /purplestore

# DEPLOY
FROM alpine:3.17

WORKDIR /app

RUN addgroup -g 1001 -S runner && \
    adduser -u 1001 -S runner -G runner

USER runner:runner

# copy purplestore from builder
COPY --from=builder /purplestore /purplestore

EXPOSE 8080

ENTRYPOINT [ "/purplestore" ]
