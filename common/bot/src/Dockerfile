FROM golang:1.7
MAINTAINER FlexConstructor <flexconstructor@gmail.com>
RUN mkdir /opt/video                                  \
    && chmod -644 /opt/video                          \
    && wget -O /opt/video/sample.flv \
    http://techslides.com/demos/sample-videos/small.flv

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY main.go /go/src/app
COPY github.com /go/src/github.com
COPY gopkg.in /go/src/gopkg.in
RUN go-wrapper install
CMD ["go-wrapper", "run","-flv_file","/opt/video/sample.flv","-redis","localhost:6379"]
