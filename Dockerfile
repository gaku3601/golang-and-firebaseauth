FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/server
COPY . .
RUN go build main.go

# runtime image
FROM golang:1.13.1-alpine3.10
RUN set -ex \
    && apk add --no-cache --virtual build-dependencies \
    build-base \
    git \
    && go get -ldflags "-extldflags -static" bitbucket.org/liamstask/goose/cmd/goose \
    && apk del build-dependencies \
    && apk add --no-cache mysql-client
COPY --from=builder /go/src/server /app

WORKDIR /app
CMD sh start.sh