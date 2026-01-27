package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"id": id, "status": "ok"})



}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/", getUserHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
    
}
