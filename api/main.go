package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	// Register the routes and handlers
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("jopa"))
	})
	// Run the server
	http.ListenAndServe(":8080", mux)
}
