FROM easypi/alpine-arm
MAINTAINER Vincent Vieira <vincent.vieira@supinfo.com>

RUN apk add --no-cache go git gcc
ENV GOPATH=/home/docker
ENV PROJECTPATH=${GOPATH}/src/vieira.io
ENV GOROOT=/usr/lib/go

RUN mkdir -p ${GOPATH}
ADD src ${GOPATH}/src/

WORKDIR ${PROJECTPATH}
RUN go get && go test ./... && go build -i -o ../../bin/collector vieira.io

RUN apk del git gcc

ENTRYPOINT ["${GOPATH}/bin/collector"]
