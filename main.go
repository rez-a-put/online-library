package main

import (
	"log"
	"net/http"
	"online-library/handler"
	"online-library/utils"

	"github.com/gorilla/mux"
)

var r *mux.Router

func init() {
	r = mux.NewRouter()
}

func main() {
	r.HandleFunc("/library", handler.GetList).Methods("GET")
	r.HandleFunc("/library/set_pickup", handler.SetPickup).Methods("POST")

	log.Fatal(http.ListenAndServe(utils.GetEnvByKey("SERVER_HOST"), r))
}
