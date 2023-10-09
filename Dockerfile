FROM golang:1.21-alpine3.18 as builder

WORKDIR /build

RUN apk add -U --no-cache ca-certificates

COPY . /build/
RUN go build


FROM alpine:3.18
EXPOSE 3000

ENV GODEBUG netdns=go

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/drone-multipipeline-abort /bin/

CMD ["/bin/drone-multipipeline-abort"]
