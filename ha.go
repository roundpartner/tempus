package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Check(router *mux.Router) {
	check := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
	metrics := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}
	router.HandleFunc("/check", check).Methods("GET")
	router.HandleFunc("/metrics", metrics).Methods("GET")
}
