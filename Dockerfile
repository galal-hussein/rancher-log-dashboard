FROM alpine:3.4
MAINTAINER Hussein Galal hussein@rancher.com

EXPOSE 5000

ADD . /app
WORKDIR /app

ENTRYPOINT ["/app/rancher-log-dashboard"]
