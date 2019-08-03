FROM golang:1.11-alpine as builder

WORKDIR /go/src/github/koneko096/blacktube-graphql

RUN apk --no-cache add git dep
ADD . /go/src/github/koneko096/blacktube-graphql

RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -o bt ./server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github/koneko096/blacktube-graphql/bt .

EXPOSE 8080
CMD ["./bt"]
