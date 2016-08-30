FROM alpine:3.4
MAINTAINER carumisu9

RUN apk add --no-cache --update --virtual=build-deps go  \
	&& apk del build-deps

COPY ./bin/kubedeploy /

ENTRYPOINT ["/kubedeploy"]
