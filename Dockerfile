FROM alpine:3.4
MAINTAINER carumisu9

ENV KUBECTL_VERSION v1.3.3
ENV GOPATH /go

COPY . /go/src/github.com/carumisu9/kubedeploy

RUN apk add --no-cache --update --virtual=build-deps curl git go make  \
  && cd /go/src/github.com/carumisu9/ \
  && go get -v ./... \
  && make \
  && cp bin/kubedeploy /kubedeploy \
  && wget -qO /usr/local/bin/kubectl "https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl" \
  && chmod +x /usr/local/bin/kubectl \
  && cd / \
  && apk del --purge wget \
  && apk del build-deps \
  && rm /var/cache/apk/* \
  && rm -rf /go

ENTRYPOINT ["/kubedeploy"]
