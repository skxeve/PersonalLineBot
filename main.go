package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", Index)
	router.Get("/number/{id}", Number)

	router.NotFound(CustomNotFound)

	log.Fatal(http.ListenAndServe(":8080", router))
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
