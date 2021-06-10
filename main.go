package main

import (
	"demo/sources"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Home_function(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dms := sources.Alienvault_enum()

	var dom1 []string

	for index, _ := range dms {
		dom1 = append(dom1, dms[index].Hostname)
	}

	json.NewEncoder(w).Encode(dom1)
}

func Subdomain_enum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dms1 := sources.Anubis_enum()

	json.NewEncoder(w).Encode(dms1)
}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", Home_function).Methods("GET")

	mux.HandleFunc("/subdomain", Subdomain_enum).Methods("GET")

	http.ListenAndServe(":8000", mux)

}
