FROM golang:1.13-alpine3.10

ENV WIMDIR ${GOPATH}/src/github.com/wim
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io 

VOLUME ${WIMDIR}
WORKDIR ${WIMDIR}

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g" /etc/apk/repositories \
&& apk update \
&& apk add git \
&& go get -u github.com/kardianos/govendor

