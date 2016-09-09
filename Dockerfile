FROM alpine:3.4
MAINTAINER carumisu9

ENV KUBECTL_VERSION v1.3.3
ENV GOPATH /go

COPY . /go/src/github.com/carumisu9/kubedeploy

RUN apk add --no-cache ca-certificates openssl
RUN apk add --no-cache --update --virtual=build-deps curl git go make  \
  && cd /go/src/github.com/carumisu9/kubedeploy \
  && go get -v ./... \
  && make \
  && cp bin/kubedeploy /kubedeploy \
  && cd / \
  && apk del build-deps \
  && rm /var/cache/apk/* \
  && rm -rf /go

ENTRYPOINT ["/kubedeploy"]
