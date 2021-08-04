ARG GO_VERSION=1.15.6

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /go/src/github/ProjectBlacktube/blacktube-graphql

RUN apk --no-cache add git

ADD . /go/src/github/ProjectBlacktube/blacktube-graphql
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/bt ./server


FROM gcr.io/distroless/static AS final

USER nonroot

COPY --from=builder --chown=nonroot:nonroot /go/bin/bt /app
COPY config /config

EXPOSE 8080
ENTRYPOINT ["/app"]
