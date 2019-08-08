FROM golang:1.12-alpine as builder

WORKDIR /go/src/github/koneko096/blacktube-graphql

RUN apk --no-cache add git

# Get latest dep
RUN go get -u github.com/golang/dep/cmd/dep

ADD . /go/src/github/koneko096/blacktube-graphql
RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/bt ./server


FROM alpine:latest

WORKDIR /root

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/bt .
COPY config ./config

EXPOSE 8080
CMD ["./bt"]
