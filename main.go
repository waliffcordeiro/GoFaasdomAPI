package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	function "github.com/waliffcordeiro/GoFaasdomAPI/functions"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/latency", function.Go_latency).Methods("GET")
	router.HandleFunc("/matrix", function.Go_matrix).Methods("GET")
	router.HandleFunc("/filesystem", function.Go_filesystem).Methods("GET")
	router.HandleFunc("/factors", function.Go_factors).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
