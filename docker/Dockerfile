FROM alpine
RUN mkdir /lib64 && ln -s /lib/ld-musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk add --no-cache ca-certificates
ENTRYPOINT ["/app/bin/mpa"]
