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

// Return all memories
func allMemoryHandler(w http.ResponseWriter, h *http.Request) {
	brain := getBrain()
	returnJson(brain, w, h)
}

// Return a random memory
func randomMemoryHandler(w http.ResponseWriter, h *http.Request) {
	brain := getBrain()
	returnJson(brain, w, h)
}

// Return a specific memory
func getMemoryHandler(w http.ResponseWriter, h *http.Request) {
	brain := getBrain()
	returnJson(brain, w, h)
}

// Add a memory
func addMemoryHandler(w http.ResponseWriter, h *http.Request) {
	brain := getBrain()
	returnJson(brain, w, h)
}

// Change a memory
func changeMemoryHandler(w http.ResponseWriter, h *http.Request) {
	brain := getBrain()
	returnJson(brain, w, h)
}

// Remove a memory
func removeMemoryHandler(w http.ResponseWriter, h *http.Request) {
	brain := getBrain()
	returnJson(brain, w, h)
}
