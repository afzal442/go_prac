FROM golang:1.16-alpine

# FROM golang:latest

RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /web4api

EXPOSE 5000

CMD [ "/web4api" ] 