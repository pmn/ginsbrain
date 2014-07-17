package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := r.RemoteAddr
		if len(remoteAddr) == 0 {
			remoteAddr = r.Header.Get("x-forwarded-for")
		}
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

var brain Brain

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", randomMemoryHandler).Methods("GET")
	r.HandleFunc("/memories/random", randomMemoryHandler).Methods("GET")
	r.HandleFunc("/memories", allMemoryHandler).Methods("GET")
	r.HandleFunc("/memories/{id}", getMemoryHandler).Methods("GET")
	r.HandleFunc("/memories", addMemoryHandler).Methods("POST")
	r.HandleFunc("/memories/{id}", changeMemoryHandler).Methods("PUT")
	r.HandleFunc("/memories/{id}", removeMemoryHandler).Methods("DELETE")
	r.HandleFunc("/memories/search/{term}", searchMemoryHandler).Methods("GET")

	http.Handle("/", r)

	log.Printf("[+] Uh, hi! My brain is running on port %s", port)
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
}

func init() {
	brain.Load()
}
