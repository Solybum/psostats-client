# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY . ./

RUN go mod download

ENV AWS_ACCESS_KEY_ID=
ENV AWS_SECRET_ACCESS_KEY=
ENV AWS_DEFAULT_REGION=

EXPOSE 80

ENTRYPOINT ["go", "run", "/app/server/cmd/main.go"]
