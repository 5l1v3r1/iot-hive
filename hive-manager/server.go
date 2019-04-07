package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createHive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func hiveStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/init-hive", createHive).Methods("POST")
	r.HandleFunc("/api/hive-status", hiveStatus).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}
