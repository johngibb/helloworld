ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY . .
RUN go build -v -o /helloworld .


FROM debian:bookworm

COPY --from=builder /helloworld /usr/local/bin/
CMD ["helloworld"]
