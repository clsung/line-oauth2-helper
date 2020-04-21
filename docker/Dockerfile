FROM golang:1.14-alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache ca-certificates bash git
RUN go get github.com/clsung/line-oauth2-helper/cmd/ljwt_server && \
    go get github.com/clsung/line-oauth2-helper/cmd/line_jwt

FROM alpine:3.11

EXPOSE 8080
COPY --from=builder /go/bin/line_jwt /usr/bin/line_jwt
COPY --from=builder /go/bin/ljwt_server /usr/bin/ljwt_server
COPY --from=builder /go/src/github.com/clsung/line-oauth2-helper/cmd/ljwt_server/static/ /usr/bin/static/
WORKDIR /usr/bin/

CMD ljwt_server
