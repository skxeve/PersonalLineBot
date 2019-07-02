package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", Index)
	router.Get("/number/{id}", Number)

	router.NotFound(CustomNotFound)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello HttpRouter Index")
}

func Number(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	fmt.Fprintf(w, "Number parameter is %s", paramID)
}

func CustomNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Custom 404 page.")
}
