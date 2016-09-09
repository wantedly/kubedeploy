FROM alpine:3.4

ENV GOPATH /go

COPY . /go/src/github.com/wantedly/kubedeploy

RUN apk add --no-cache ca-certificates openssl
RUN apk add --no-cache --update --virtual=build-deps curl git go make  \
  && cd /go/src/github.com/wantedly/kubedeploy \
  && go get -v ./... \
  && make \
  && cp bin/kubedeploy /kubedeploy \
  && cd / \
  && apk del build-deps \
  && rm /var/cache/apk/* \
  && rm -rf /go

ENTRYPOINT ["/kubedeploy"]
