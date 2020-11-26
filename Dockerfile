FROM golang:1.10-alpine3.7 as build

WORKDIR /go/src/hw

COPY hw.go /go/src/hw/

RUN apk add git 
RUN go get github.com/gorilla/mux && go build -o hw

#RUN pwd && ls
RUN cp  hw /usr/local/bin/hw

ENTRYPOINT ["/usr/local/bin/hw"]
