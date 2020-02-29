FROM golang:1.13-alpine as builder

WORKDIR /go/src/github/ProjectBlacktube/blacktube-graphql

RUN apk --no-cache add git

ADD . /go/src/github/ProjectBlacktube/blacktube-graphql
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/bt ./server


FROM alpine:latest

WORKDIR /root

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/bt .
COPY config ./config

EXPOSE 8080
CMD ["./bt"]
