FROM golang:1.9

ADD ./ /go/src/github.com/kvss/simple-charity-server
WORKDIR /go/src/github.com/kvss/simple-charity-server

RUN go get -u github.com/golang/dep/cmd/dep && go get -u -v github.com/go-task/task/cmd/task && dep ensure && go build && go-wrapper install
CMD ["go-wrapper", "run"]