FROM resin/rpi-raspbian:jessie
MAINTAINER Vincent Vieira <vincent.vieira@supinfo.com>

#Installing required stuff
RUN apt-get update && apt-get install -y git build-essential wget

#Installing Go runtime
RUN wget -O go.tar.gz https://storage.googleapis.com/golang/go1.7.5.linux-armv6l.tar.gz && sudo tar -C /usr/local -xzf go.tar.gz
RUN rm go.tar.gz
ENV PATH /usr/local/go/bin:$PATH

#Installing bcm2835 library
RUN wget -O bcm2835.tar.gz http://www.airspayce.com/mikem/bcm2835/bcm2835-1.52.tar.gz && tar zxvf bcm2835.tar.gz
WORKDIR bcm2835-1.52
RUN ./configure && make && sudo make install

ENV GOPATH=/home/docker
ENV PROJECTPATH=${GOPATH}/src/
ENV GOROOT=/usr/local/go

RUN mkdir -p ${PROJECTPATH}
ADD src ${PROJECTPATH}

WORKDIR ${PROJECTPATH}
RUN go get -t ./... && go test vieira.io/... && go build -i -o ../bin/collector vieira.io

ENTRYPOINT sudo $GOPATH/bin/collector
