package main

import (
	"fmt"
	"github.com/go-chi/chi"
	//h "github.com/skxeve/PersonalLineBot/line/http"
	gcp_logging "cloud.google.com/go/logging"
	"context"
	"github.com/skxeve/PersonalLineBot/line/log"
	l "log"
	"net/http"
	"os"
	"strings"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", Index)
	router.Get("/sample", Sample)
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
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello?")
}

func Sample(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := gcp_logging.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		// Handle error.
		l.Fatal("logger client open error: %v", err)
	}
	// Initialize a logger
	lg := client.Logger("appengine.googleapis.com/applog")

	// Add entry to log buffer
	lg.Log(gcp_logging.Entry{
		Payload:  "something happened!",
		Severity: gcp_logging.Critical,
	})
	err = lg.LogSync(ctx, gcp_logging.Entry{
		Payload:  "something sync happened!",
		Severity: gcp_logging.Critical,
	})
	if err != nil {
		l.Printf("logsync failed: %v", err)
	}
	err = client.Close()
	if err != nil {
		// TODO: Handle error.
		l.Fatal("logger client close error: %v", err)
	}
	l.Printf("[LOG] Sample GAE logging test normal log output: %v", r.Header)
	fmt.Fprintf(w, "Sample GAE logging with golang %v", r.Header)
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
