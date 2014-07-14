package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Return the data in JSON format. This is the default return method.
func returnJson(obj interface{}, w http.ResponseWriter, h *http.Request) {
	// Don't cache json returns. This is to work around ie's weird caching behavior
	w.Header().Set("Cache-Control", "no-cache")
	// Set the content type to json
	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprint(w, string(j))
}

func defaultHandler(w http.ResponseWriter, h *http.Request) {
  fmt.Fprint(w, "cool, running")
}
