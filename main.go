package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/waliffcordeiro/GoFaasdomAPI/functions"
)

func main() {
	router := mux.NewRouter()
    router.HandleFunc("/latency", functions.Go_latency).Methods("GET")
	router.HandleFunc("/matrix", functions.Go_matrix).Methods("GET")
	router.HandleFunc("/filesystem", functions.Go_filesystem).Methods("GET")
	router.HandleFunc("/factors", functions.Go_factors).Methods("GET")
	http.ListenAndServe(":8000", router)
}