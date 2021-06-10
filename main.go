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

	json.NewEncoder(w).Encode(dms)
}

func Type_function(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dms1 := sources.Alienvault_enum()
	params := mux.Vars(r)

	for index, _ := range dms1 {
		if dms1[index].Hostname != "" && dms1[index].Record_type == params["type"] {
			json.NewEncoder(w).Encode(dms1)
		}
	}

}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", Home_function).Methods("GET")
	mux.HandleFunc("/{type}", Type_function).Methods("GET")

	http.ListenAndServe(":8000", mux)

}
