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
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler).Methods("GET")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.Handle("/", r)
	r.HandleFunc("/memory/random", randomMemoryHandler).Methods("GET")
	r.HandleFunc("/memory/all", allMemoryHandler).Methods("GET")
	r.HandleFunc("/memory/[0-9]", getMemoryHandler).Methods("GET")
	r.HandleFunc("/memory", addMemoryHandler).Methods("POST")
	r.HandleFunc("/memory/[0-9]", changeMemoryHandler).Methods("PUT, PATCH")
	r.HandleFunc("/memory/[0-9]", removeMemoryHandler).Methods("DELETE")

	log.Printf("[+] Uh, hi! My brain is running on port %s", port)
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
}

func init() {
	brain.Load()
}
