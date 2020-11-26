package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// commenta
var port = "8080"

func main() {
	log.Fatal(http.ListenAndServe(":"+port, router()))
}

func router() http.Handler {
	r := mux.NewRouter()
	r.Path("/greeting").Methods(http.MethodGet).HandlerFunc(greet)
	r.Path("/healthcheck").Methods(http.MethodGet).HandlerFunc(healthcheck)
	return r
}

func healthcheck(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("ok"))
}

func greet(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("Hello, world!"))
}
