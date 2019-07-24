package main

import (
	"fmt"
	"github.com/go-chi/chi"
	h "github.com/skxeve/PersonalLineBot/line/http"
	l "log"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		router.ServeHTTP(w, r)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	l.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello?")
}

func EnvList(w http.ResponseWriter, r *http.Request) {
	c := h.NewContext(r)
	c.Logger.Infof("envlist %s", strings.Join(os.Environ(), ", "))
	fmt.Fprintf(w, strings.Join(os.Environ(), "\n"))
}

func LineWebHook(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")
	c := h.NewContext(r)
	c.Logger.Infof("WebHook Received %s", accountId)
	fmt.Fprintf(w, "WebHook Received %s", accountId)
}

func CustomNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "404")
}
