FROM alpine:3.20
COPY zapwerx-cli /usr/bin/zapwerx-cli
ENTRYPOINT ["/usr/bin/zapwerx-cli"]