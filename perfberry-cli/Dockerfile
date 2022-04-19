FROM golang:1.12.9-alpine3.10 AS build

RUN mkdir -p $GOPATH/src/github.com/perfberry/cli
WORKDIR $GOPATH/src/github.com/perfberry/cli

ADD . .
RUN CGO_ENABLED=0 go build -o /usr/local/bin/perfberry-cli -ldflags '-w'


FROM alpine:3.14

RUN apk add --update ca-certificates && rm /var/cache/apk/*

COPY --from=build /usr/local/bin/perfberry-cli /usr/local/bin/

CMD perfberry-cli
