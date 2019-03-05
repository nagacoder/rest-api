package handlers

import (
	"encoding/json"
	"net/http"
)

// handle the root route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	values := map[string]string{"message": "Server is Running"}
	result, _ := json.Marshal(values)
	w.Write(result)
}
