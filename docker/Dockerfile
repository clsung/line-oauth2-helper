FROM golang:1.26-alpine AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/line_jwt ./cmd/line_jwt

WORKDIR /src/cmd/ljwt_server
RUN go mod download && go build -o /go/bin/ljwt_server .

FROM alpine:3.11

EXPOSE 8080
COPY --from=builder /go/bin/line_jwt /usr/bin/line_jwt
COPY --from=builder /go/bin/ljwt_server /usr/bin/ljwt_server
COPY --from=builder /src/cmd/ljwt_server/static/ /usr/bin/static/
WORKDIR /usr/bin/

CMD ljwt_server
