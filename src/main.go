package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/insights", GetInsights).Methods("GET")
	fmt.Println("Server is listening on port 8000...")
	http.ListenAndServe(":8000", router)
}

func GetInsights(w http.ResponseWriter, r *http.Request) {
	insights := []string{ "Total Users: 1000", "Total Sessions: 5000" }
	json.NewEncoder(w).Encode(insights)
}