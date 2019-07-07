package main

import (
	"./line/log"
	"fmt"
	"github.com/go-chi/chi"
	//l "log"
	"net/http"
	"os"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", Index)
	router.Post("/webhook/{id}", LineWebHook)

	router.NotFound(CustomNotFound)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
	//l.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func Index(w http.ResponseWriter, r *http.Request) {
	logger := log.Logger{}
	logger.Debugf("Index request.")
	fmt.Fprintf(w, "Hello?")
}

func LineWebHook(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")
	logger := log.Logger{}
	logger.Infof("WebHook Received %s", accountId)
	fmt.Fprintf(w, "WebHook Received %s", accountId)
}

func CustomNotFound(w http.ResponseWriter, r *http.Request) {
	logger := log.Logger{}
	logger.Warningf("NotFound request.")
	fmt.Fprintf(w, "404")
}
