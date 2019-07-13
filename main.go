package main

import (
	"fmt"
	"github.com/go-chi/chi"
	//h "github.com/skxeve/PersonalLineBot/line/http"
	"github.com/skxeve/PersonalLineBot/line/log"
	"google.golang.org/appengine"
	gaelog "google.golang.org/appengine/log"
	"net/http"
	"os"
	"reflect"
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
	gae_instance := os.Getenv("GAE_INSTANCE")
	is_gae_env := gae_instance != ""
	if is_gae_env {
		appengine.Main()
	} else {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	//c := h.GetHttpContext(r)
	//c.Logger.Debugf("Index request.")
	ctx := appengine.NewContext(r)
	logger := log.Logger{}
	logger.Infof("r:%s ctx:%s", reflect.TypeOf(r), reflect.TypeOf(ctx))
	gaelog.Debugf(ctx, "DebugPrint1")
	gaelog.Infof(ctx, "InfoPrint2")
	gaelog.Warningf(ctx, "WarnPrint3")
	gaelog.Errorf(ctx, "ERrorPrint4")
	fmt.Fprintf(w, "Hello?")
}

func Sample(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	query := &gaelog.Query{
		AppLogs:  true,
		Versions: []string{"1"},
	}
	for results := query.Run(ctx); ; {
		record, err := results.Next()
		if err == gaelog.Done {
			gaelog.Infof(ctx, "Done processing results")
			break
		}
		if err != nil {
			gaelog.Errorf(ctx, "Failed to retrieve next log: %v", err)
			break
		}
		gaelog.Infof(ctx, "Saw record %v", record)
		fmt.Fprintf(w, "Saw record %v", record)
	}
	gaelog.Debugf(ctx, "Done for-loop in Sample")
	fmt.Fprintf(w, "Sample GAE logging with golang")
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
