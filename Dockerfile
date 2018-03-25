FROM golang:latest
WORKDIR /go/src/github.com/m-nakamura145/golang-docker-compose-sample
ADD . /go/src/github.com/m-nakamura145/golang-docker-compose-sample
EXPOSE 5000
