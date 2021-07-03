FROM golang:alpine

RUN set -ex; apk update; apk add --no-cache git; apk add bash
WORKDIR /go/src/github.com/Projetos/go_boilerplate

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o go_boilerplate

EXPOSE 3000/tcp
CMD ["./go_boilerplate"]
