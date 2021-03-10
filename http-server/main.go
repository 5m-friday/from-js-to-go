package main

import (
	"log"
	"net/http"
)

func main() {
	// KEEP IN MIND, ONLY USING STD LIB
	// Production ready HTTP Server
	http.HandleFunc("/", 世界)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func 世界(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "OK"}`))
}
