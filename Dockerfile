FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/rarimo/galxe-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/galxe-svc /go/src/github.com/rarimo/galxe-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/galxe-svc /usr/local/bin/galxe-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["galxe-svc"]
