FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -v
RUN go get github.com/lib/pq
RUN go get gopkg.in/Nhanderu/brdoc.v1
RUN go build