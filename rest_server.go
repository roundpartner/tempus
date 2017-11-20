package main

import "github.com/gorilla/mux"
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func ListenAndServe() {
	rs := NewRestServer()
	server := &http.Server{Addr: ":7373", Handler: rs.Router}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		<-c
		signal.Stop(c)
		fmt.Println("http: Server shutting down gracefully")
		server.Shutdown(nil)
	}()

	fmt.Println("http: Server starting")
	err := server.ListenAndServe()
	if nil != err {
		fmt.Println(err.Error())
	}
}

type RestServer struct {
	Router    *mux.Router
	Store     *Store
	Generator *TokenGenerator
}

func NewRestServer() *RestServer {
	rs := &RestServer{}
	rs.Router = mux.NewRouter()
	rs.Router.HandleFunc("/", rs.AddToken).Methods("POST")
	rs.Router.HandleFunc("/{user_id}/{token}", rs.GetToken).Methods("GET")

	rs.Store = New()
	rs.Generator = NewTokenGenerator()
	return rs
}

func (rs *RestServer) AddToken(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	token := &Token{}
	err := decoder.Decode(token)
	if err != nil {
		log.Printf("Unable to decode request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token = rs.Generator.Get(token.User, token.Scenario)

	rs.Store.Add(token, time.Hour*24*3)

	data, _ := token.MarshalBinary()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (rs *RestServer) GetToken(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId, err := strconv.ParseInt(params["user_id"], 10, 64)
	if err != nil {
		log.Printf("Unable to decode request: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	token, err := rs.Store.Get(params["token"])
	if err != nil {
		log.Printf("Unable to get token: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if token == nil {
		log.Printf("Token not found: %d/%s\n", userId, params["token"])
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if userId != token.User {
		log.Printf("User does not match: %d != %d\n", userId, token.User)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := token.MarshalBinary()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}