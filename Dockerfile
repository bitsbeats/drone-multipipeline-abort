FROM alpine:3.18 as alpine
RUN apk add -U --no-cache ca-certificates

FROM alpine:3.18
EXPOSE 3000

ENV GODEBUG netdns=go

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ADD drone-multipipeline-abort /bin/
ENTRYPOINT ["/bin/drone-multipipeline-abort"]