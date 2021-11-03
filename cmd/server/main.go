package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

var db = map[string]string{
	"gui": "nat",
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/", handleKey)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHealth(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"name":    "clio-server",
		"version": "v0.0.1",
	}

	data, err := json.Marshal(response)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}

func handleKey(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost, http.MethodPut:
		//TODO
	case http.MethodGet:
		trimmedPath := strings.TrimPrefix(r.URL.Path, "/kv")
		if trimmedPath == r.URL.Path {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		key := strings.TrimPrefix(trimmedPath, "/")

		value, ok := db[key]
		if !ok {
			value = ""
		}

		response := map[string]string{
			"value": value,
		}

		err := json.NewEncoder(rw).Encode(response)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
