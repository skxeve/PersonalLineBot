FROM golang:1.11

RUN go get -u github.com/go-chi/chi

WORKDIR /go/src/github.com/skxeve/PersonalLineBot
CMD go run main.go
