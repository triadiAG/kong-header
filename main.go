package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":9090"

func hello(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Auth")

	body := map[string]interface{}{
		"header": header,
	}

	data, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/backoffice/coba", hello).Methods("GET")

	fmt.Println("Running on port", PORT)
	http.ListenAndServe(PORT, r)
}
