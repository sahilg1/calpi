FROM golang

ADD  . /go/src/calpi
RUN cd /go/src/calpi &&\
    go build &&\
    go install
CMD calpi
