FROM golang

ADD . /go/src/github.com/transcranial/tricorder

RUN cd /go/src/github.com/transcranial/tricorder \
    && go get -v \
    && go build tricorder.go