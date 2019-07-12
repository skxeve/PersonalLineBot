package main

import (
	"fmt"
	"github.com/go-chi/chi"
	h "github.com/skxeve/PersonalLineBot/line/http"
	"github.com/skxeve/PersonalLineBot/line/log"
	"net/http"
	"os"
	"strings"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", Index)
	router.Get("/env/list", EnvList)
	router.Post("/webhook/{id}", LineWebHook)

	router.NotFound(CustomNotFound)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func Index(w http.ResponseWriter, r *http.Request) {
	c := h.GetHttpContext(r)
	c.Logger.Debugf("Index request.")
	fmt.Fprintf(w, "Hello?")
}

func EnvList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, strings.Join(os.Environ(), "\n"))
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
