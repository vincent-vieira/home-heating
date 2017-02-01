FROM easypi/alpine-arm
MAINTAINER Vincent Vieira <vincent.vieira@supinfo.com>

RUN apk add --no-cache go git
ENV GOPATH=/home/docker
ENV PROJECTPATH=${GOPATH}/src/
ENV GOROOT=/usr/lib/go

RUN mkdir -p ${PROJECTPATH}
ADD src ${PROJECTPATH}

WORKDIR ${PROJECTPATH}
RUN go get -t ./... && go test vieira.io/... && go build -i -o ../bin/collector vieira.io

RUN apk del git

ENTRYPOINT ["${GOPATH}/bin/collector"]
