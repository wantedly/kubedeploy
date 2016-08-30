FROM alpine:3.4
MAINTAINER carumisu9

ENV GOPATH /go

COPY . /go/src/github.com/carumisu9/kubedeploy
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache --update --virtual=build-deps curl git go make  \
		&& cd /go/src/github.com/carumisu9/kubedeploy \
		&& go get -v ./... \
		&& make \
		&& cp bin/kubedeploy /kubedeploy \
		&& cd / \
		&& rm -rf /go \
		&& apk del build-deps

ENTRYPOINT ["/kubedeploy"]
