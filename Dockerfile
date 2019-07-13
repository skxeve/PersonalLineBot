FROM golang:1.12

#RUN go get -u github.com/go-chi/chi
#RUN go get -u google.golang.org/appengine/log
#RUN go get -u google.golang.org/appengine


WORKDIR /go/src/github.com/skxeve/PersonalLineBot
ENV GO111MODULE on
RUN go mod download
CMD go run main.go
