FROM golang:1.11

RUN go get -u github.com/go-chi/chi
RUN go get -u google.golang.org/appengine/log
RUN go get -u google.golang.org/appengine

WORKDIR /go/src/github.com/skxeve/PersonalLineBot
CMD go run main.go
