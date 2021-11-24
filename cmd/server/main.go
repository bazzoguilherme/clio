package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var db = map[string]string{
	"gui": "nat",
}

type bodyRequest struct {
	value string
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
		key := getKey(*r.URL)
		if key == "" {
			fmt.Print("error key")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		body := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		value := body["value"]

		db[key] = value

	case http.MethodGet:
		key := getKey(*r.URL)
		if key == "" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

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

func getKey(reqURL url.URL) string {
	trimmedPath := strings.TrimPrefix(reqURL.Path, "/kv")
	if trimmedPath == reqURL.Path {
		return ""
	}
	return strings.TrimPrefix(trimmedPath, "/")
}
