FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

RUN cd /build/GO_PRAC/REST-API && go build

EXPOSE 8080

ENTRYPOINT [ "/build/GO_PRAC/REST-API/routes" ] 