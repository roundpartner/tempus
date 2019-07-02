package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func ListenAndServe(port int) {
	address := fmt.Sprintf(":%d", port)

	rs := NewRestServer()
	server := &http.Server{Addr: address, Handler: rs.Router}
	ShutdownGracefully(server)

	log.Printf("Server starting on port %d\n", port)
	err := server.ListenAndServe()
	if nil != err {
		log.Printf("[INFO] [%s] %s", ServiceName, err.Error())
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
	Check(rs.Router)
	rs.Router.HandleFunc("/", rs.AddToken).Methods("POST")
	rs.Router.HandleFunc("/{user_id}/{token}", rs.GetToken).Methods("GET")
	rs.Router.HandleFunc("/{user_id}/{scenario}/{token}", rs.GetToken).Methods("GET")

	rs.Store = New()
	rs.Generator = NewTokenGenerator()
	return rs
}

func (rs *RestServer) AddToken(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	requestToken := &Token{}
	err := decoder.Decode(requestToken)
	if err != nil {
		log.Printf("Unable to decode request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Adding %s token for user %d\n", requestToken.Scenario, requestToken.User)
	responseToken := rs.Generator.Get(requestToken.User, requestToken.Scenario)
	responseToken.Meta = requestToken.Meta

	rs.Store.AddLater(responseToken, responseToken.Expires())

	data, _ := responseToken.MarshalBinary()

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

	validator := UserValidator(userId)
	if params["scenario"] != "" {
		validator = UserScenarioValidator(userId, params["scenario"])
	}
	log.Printf("Looking up token for user %d\n", userId)
	token, err := rs.Store.Get(params["token"], validator)
	if err != nil {
		log.Printf("Unable to get token: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if token == nil {
		if params["scenario"] != "" {
			log.Printf("token not found: %d/%s/%s\n", userId, params["scenario"], params["token"])
		} else {
			log.Printf("token not found: %d/%s\n", userId, params["token"])
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if userId != token.User {
		log.Printf("User does not match: %d != %d\n", userId, token.User)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := token.MarshalBinary()
	if err != nil {
		log.Printf("Unable to marshal token: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
