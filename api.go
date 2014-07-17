package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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
	returnJson(brain, w, h)
}

// Return a random memory
func randomMemoryHandler(w http.ResponseWriter, h *http.Request) {
	random_id := rand.Intn(brain.GetNextId() - 1)
	m := brain.Memories[random_id]
	returnJson(m, w, h)
}

// Return a specific memory
func getMemoryHandler(w http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id_string := vars["id"]

	id, err := strconv.Atoi(id_string)
	if err != nil {
		fmt.Fprint(w, "Could not parse Id")
	}

	for _, v := range brain.Memories {
		if v.Id == id {
			returnJson(v, w, h)
		}
	}
}

// Add a memory
func addMemoryHandler(w http.ResponseWriter, h *http.Request) {
	var m Memory
	b := json.NewDecoder(h.Body)
	b.Decode(&m)

	memory := brain.Add(m)
	returnJson(memory, w, h)
}

func searchMemoryHandler(w http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	term := vars["term"]

	var results []Memory

	for _, v := range brain.Memories {
		if strings.Contains(v.Text, term) {
			results = append(results, v)
		}
	}

	returnJson(results, w, h)
}

// Change a memory
func changeMemoryHandler(w http.ResponseWriter, h *http.Request) {
	var m Memory
	b := json.NewDecoder(h.Body)
	b.Decode(&m)

	// The ID should come from the URL route, not the memory object that was posted
	vars := mux.Vars(h)
	mem_id_str := vars["id"]

	mem_id, err := strconv.Atoi(mem_id_str)
	if err != nil {
		panic("ID was not valid")
	}
	m.Id = mem_id

	memory := brain.Update(m)

	returnJson(memory, w, h)
}

// Remove a memory
func removeMemoryHandler(w http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id_str := vars["id"]
	id, err := strconv.Atoi(id_str)

	if err != nil {
		panic("ID was not an integer")
	}

	// Find the item to be deleted
	var memory Memory
	for k, v := range brain.Memories {
		if v.Id == id {
			v.Active = false
			memory = brain.Memories[k]
		}
	}

	returnJson(memory, w, h)
}
