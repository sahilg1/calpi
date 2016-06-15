FROM golang

ADD  . /go/src/Users/sahgupta/Documents/gowork/src/github.com/sahilg1/Training/calpi
RUN cd /go/src/Users/sahgupta/Documents/gowork/src/github.com/sahilg1/Training/calpi &&\
    go build &&\
    go install
CMD calpi
