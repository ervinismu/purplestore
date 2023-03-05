# reference : https://docs.docker.com/build/building/multi-stage/
# BUILD
FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN touch .env

COPY . /app

RUN go build -o /purplestore

# DEPLOY
# reference : https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /purplestore /purplestore

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/purplestore" ]
